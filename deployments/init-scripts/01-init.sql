CREATE DATABASE metrics_db;
\c metrics_db

-- TimescaleDB 확장 설치
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;  -- CASCADE 추가

-- 메트릭 테이블 생성
CREATE TABLE metrics (
                         time        TIMESTAMPTZ NOT NULL,
                         metric_name TEXT NOT NULL,
                         value       DOUBLE PRECISION NOT NULL,
                         labels      JSONB
);

-- 하이퍼테이블로 변환
SELECT create_hypertable('metrics', 'time');
-- 하이퍼테이블로 변환하기 전에는 show_chunks 를 사용할 수 없다

-- 인덱스 생성
CREATE INDEX idx_metrics_metric_name ON metrics (metric_name, time DESC);
CREATE INDEX idx_metrics_labels ON metrics USING GIN (labels);