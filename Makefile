.PHONY: run-admin-svc
run-admin-svc:
	cd apps/admin-svc && go run main.go

.PHONY: run-setting-svc
run-setting-svc:
	cd apps/setting-svc && go run main.go

.PHONY: run-auth-svc
run-auth-svc:
	cd apps/auth-svc && go run main.go

.PHONY: run-orchestrator-svc
run-orchestrator-svc:
	cd apps/orchestrator-svc && go run main.go

.PHONY: run-time-svc
run-time-svc:
	cd apps/time-svc && go run main.go

.PHONY: run-workers
run-workers:
	cd workers/sync-worker && go run main.go