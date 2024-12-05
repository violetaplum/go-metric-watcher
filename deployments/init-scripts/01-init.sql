-- TimescaleDB 확장 활성화
CREATE EXTENSION IF NOT EXISTS timescaledb;

-- 알림 상태를 위한 enum 타입
CREATE TYPE alert_status AS ENUM ('triggered', 'resolved', 'acknowledged');

-- 알림 히스토리 테이블 생성
CREATE TABLE alert_histories (
                                 time                TIMESTAMPTZ NOT NULL,
                                 alert_rule_id      BIGINT NOT NULL,
                                 metric_name        VARCHAR(100) NOT NULL,
                                 metric_value       DOUBLE PRECISION NOT NULL,
                                 threshold_value    DOUBLE PRECISION NOT NULL,
                                 status            alert_status NOT NULL,
                                 description       TEXT,
                                 resolved_at       TIMESTAMPTZ,
                                 acknowledged_at   TIMESTAMPTZ,
                                 acknowledged_by   VARCHAR(100),
                                 target_system    VARCHAR(100) NOT NULL,
                                 severity         VARCHAR(20) NOT NULL
);

-- 하이퍼테이블로 변환 (1주일 단위로 청크 생성)
SELECT create_hypertable('alert_histories', 'time',
                         chunk_time_interval => INTERVAL '1 week');

-- 인덱스 생성
CREATE INDEX idx_alert_histories_alert_rule_id ON alert_histories(alert_rule_id, time DESC);
CREATE INDEX idx_alert_histories_status ON alert_histories(status, time DESC);
CREATE INDEX idx_alert_histories_target_system ON alert_histories(target_system, time DESC);

-- 압축 정책 설정 (optional)
ALTER TABLE alert_histories SET (
    timescaledb.compress,
    timescaledb.compress_segmentby = 'alert_rule_id,status,target_system'
    );

-- 30일 이후의 데이터는 자동 압축
SELECT add_compression_policy('alert_histories', INTERVAL '30 days');

-- 1년 이상된 데이터 자동 삭제
SELECT add_retention_policy('alert_histories', INTERVAL '1 year');