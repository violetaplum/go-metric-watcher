-- TimescaleDB 확장 생성
CREATE EXTENSION IF NOT EXISTS timescaledb;

-- 알림 히스토리 테이블 생성
CREATE TABLE IF NOT EXISTS alert_histories (
                                               id BIGSERIAL,
                                               time TIMESTAMPTZ NOT NULL,
                                               alert_rule_id BIGINT NOT NULL,
                                               metric_name TEXT NOT NULL,
                                               metric_value DECIMAL NOT NULL,
                                               threshold_value DECIMAL NOT NULL,
                                               status TEXT NOT NULL,
                                               description TEXT,
                                               resolved_at TIMESTAMPTZ,
                                               target_system TEXT NOT NULL,
                                               severity TEXT NOT NULL,
                                               created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
                                               PRIMARY KEY (id, time)
    );

-- 하이퍼테이블 생성
SELECT create_hypertable('alert_histories', 'time', if_not_exists => TRUE, migrate_data => TRUE);

-- 인덱스 생성
CREATE INDEX IF NOT EXISTS idx_alert_histories_alert_rule_id ON alert_histories(alert_rule_id, time DESC);
CREATE INDEX IF NOT EXISTS idx_alert_histories_status ON alert_histories(status, time DESC);

-- 압축 정책 설정 (선택사항)
SELECT add_compression_policy('alert_histories', INTERVAL '7 days', if_not_exists => TRUE);