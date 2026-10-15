-- La imagen oficial ejecuta este archivo al crear el volumen por primera vez.
-- created_at se completa en MySQL; cada fila representa un evento de la app.
CREATE TABLE IF NOT EXISTS app_logs (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    level VARCHAR(10) NOT NULL,
    message TEXT NOT NULL,
    INDEX idx_app_logs_created_at (created_at)
);

-- Grafana solo necesita leer los logs. El usuario demo de la app puede escribirlos.
CREATE USER IF NOT EXISTS 'grafana_reader'@'%' IDENTIFIED BY 'grafana_demo';
GRANT SELECT ON observabilidad.app_logs TO 'grafana_reader'@'%';
