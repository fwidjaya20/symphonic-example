CREATE TABLE IF NOT EXISTS "tasks" (
   "id" UUID,
   "title" VARCHAR(255) NOT NULL,
   "description" TEXT,
   "is_completed" BOOLEAN DEFAULT false,
   "is_priority" BOOLEAN DEFAULT false,
   "due_date" DATE,
   "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   "deleted_at" TIMESTAMP,
   CONSTRAINT "pk_tasks" PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS "ix_tasks_id" ON "tasks" USING btree("id");