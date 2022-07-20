package postgresql_simple_protocol_binary_format_bench_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

func mustConnect(t testing.TB, ctx context.Context) *pgx.Conn {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatal(err)
	}
	config.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	conn, err := pgx.ConnectConfig(ctx, config)
	if err != nil {
		t.Fatal(err)
	}

	return conn
}

func BenchmarkTextFormat1Row(b *testing.B) {
	ctx := context.Background()
	conn := mustConnect(b, ctx)
	defer conn.Close(ctx)

	benchmarkQuery(b, ctx, conn, 1)
}

func BenchmarkBinaryFormat1Row(b *testing.B) {
	ctx := context.Background()
	conn := mustConnect(b, ctx)
	defer conn.Close(ctx)

	_, err := conn.Exec(ctx, `set format_binary='20,21,23,1184'`)
	if err != nil {
		b.Fatal(err)
	}

	benchmarkQuery(b, ctx, conn, 1)
}

func BenchmarkTextFormat100Rows(b *testing.B) {
	ctx := context.Background()
	conn := mustConnect(b, ctx)
	defer conn.Close(ctx)

	benchmarkQuery(b, ctx, conn, 100)
}

func BenchmarkBinaryFormat100Rows(b *testing.B) {
	ctx := context.Background()
	conn := mustConnect(b, ctx)
	defer conn.Close(ctx)

	_, err := conn.Exec(ctx, `set format_binary='20,21,23,1184'`)
	if err != nil {
		b.Fatal(err)
	}

	benchmarkQuery(b, ctx, conn, 100)
}

func benchmarkQuery(b *testing.B, ctx context.Context, conn *pgx.Conn, rowCount int32) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, err := conn.Query(ctx, `select n, 'user'||n::text, now() from generate_series(100000, 100000+$1) n`, rowCount-1)
		if err != nil {
			b.Fatal(err)
		}

		for rows.Next() {
			var id int32
			var username string
			var creationTime time.Time
			err := rows.Scan(&id, &username, &creationTime)
			if err != nil {
				b.Fatal(err)
			}
		}

		if err := rows.Err(); err != nil {
			b.Fatal(err)
		}
	}
}
