package main

import "log"

// Log informativo
func logInfo(msg string) {
	log.Printf("ℹ️ %s", msg)
}

// Log de sucesso
func logSuccess(msg string) {
	log.Printf("✅ %s", msg)
}

// Log de aviso
func logWarning(msg string) {
	log.Printf("⚠️ %s", msg)
}

// Log de erro
func logError(msg string) {
	log.Printf("❌ %s", msg)
}

// Log crítico
func logCritical(msg string) {
	log.Printf("💥 %s", msg)
}

// Log de evento
func logEvent(msg string) {
	log.Printf("📣 %s", msg)
}

// Log de debug
func logDebug(msg string) {
	log.Printf("🐞 %s", msg)
}
