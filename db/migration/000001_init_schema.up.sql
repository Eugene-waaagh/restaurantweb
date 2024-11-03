CREATE TABLE "foodcatalogue" (
                                 "id" SERIAL PRIMARY KEY,
                                 "name" VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE "food" (
                        "id" SERIAL PRIMARY KEY,
                        "name" VARCHAR(100) NOT NULL,
                        "description" VARCHAR(100) NOT NULL,
                        "price" INT NOT NULL,
                        "category_id" INT NOT NULL
);

CREATE TABLE "foodcustomization" (
                                     "id" SERIAL PRIMARY KEY,
                                     "food_id" INT NOT NULL,
                                     "customization_type" VARCHAR(50) NOT NULL,
                                     "value" VARCHAR(50) NOT NULL
);

CREATE TABLE "order" (
                         "id" SERIAL PRIMARY KEY,
                         "order_date" VARCHAR(50) NOT NULL DEFAULT TO_CHAR(NOW(), 'YYYY-MM-DD HH24:MI:SS'),
                         "total_price" INT NOT NULL,
                         "status" VARCHAR(20) NOT NULL DEFAULT 'pending'  -- Replaced ENUM with VARCHAR
);

CREATE TABLE "orderitem" (
                             "id" SERIAL PRIMARY KEY,
                             "order_id" INT NOT NULL,
                             "food_id" INT NOT NULL,
                             "quantity" INT NOT NULL DEFAULT 1,
                             "item_price" INT NOT NULL
);

CREATE TABLE "ordercustomization" (
                                      "order_item_id" INT NOT NULL,
                                      "food_customization_id" INT NOT NULL,
                                      PRIMARY KEY ("order_item_id", "food_customization_id")
);

CREATE INDEX "idx_food_category" ON "food" ("category_id");

CREATE INDEX "idx_food_customization_food" ON "foodcustomization" ("food_id");

CREATE INDEX "idx_orderitem_order" ON "orderitem" ("order_id");

CREATE INDEX "idx_orderitem_food" ON "orderitem" ("food_id");

CREATE INDEX "idx_order_customization" ON "ordercustomization" ("order_item_id", "food_customization_id");

ALTER TABLE "food"
    ADD CONSTRAINT "fk_category" FOREIGN KEY ("category_id") REFERENCES "foodcatalogue" ("id") ON DELETE RESTRICT;

ALTER TABLE "foodcustomization"
    ADD CONSTRAINT "fk_food" FOREIGN KEY ("food_id") REFERENCES "food" ("id") ON DELETE CASCADE;

ALTER TABLE "orderitem"
    ADD CONSTRAINT "fk_order" FOREIGN KEY ("order_id") REFERENCES "order" ("id") ON DELETE CASCADE;

ALTER TABLE "orderitem"
    ADD CONSTRAINT "fk_food_order_item" FOREIGN KEY ("food_id") REFERENCES "food" ("id") ON DELETE CASCADE;

ALTER TABLE "ordercustomization"
    ADD CONSTRAINT "fk_order_item" FOREIGN KEY ("order_item_id") REFERENCES "orderitem" ("id") ON DELETE CASCADE;

ALTER TABLE "ordercustomization"
    ADD CONSTRAINT "fk_food_customization" FOREIGN KEY ("food_customization_id") REFERENCES "foodcustomization" ("id") ON DELETE CASCADE;
