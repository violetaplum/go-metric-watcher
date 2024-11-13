-- 기존 DB 삭제 및 생성 (필요한 경우)
DROP DATABASE IF EXISTS metrics_db;
CREATE DATABASE metrics_db;

\c metrics_db

-- TimescaleDB 확장 설치
CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;

-- 메트릭 데이터를 저장할 테이블 생성
CREATE TABLE metrics (
                         time        TIMESTAMPTZ NOT NULL,
                         metric_name TEXT NOT NULL,
                         value       DOUBLE PRECISION NOT NULL,
                         labels      JSONB
);

-- 하이퍼테이블로 변환 (시계열 데이터 최적화)
SELECT create_hypertable('metrics', 'time');

-- 인덱스 생성
CREATE INDEX idx_metrics_metric_name ON metrics (metric_name, time DESC);
CREATE INDEX idx_metrics_labels ON metrics USING GIN (labels);