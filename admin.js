const express = require("express");
const router = express.Router();
const checkRole = require("../middleware/checkRole");
const modulController = require("../controllers/modulController");

// Gunakan middleware untuk admin
router.use(checkRole("admin"));

// CRUD Modul
router.get("/modul", modulController.getAllModul);
router.post("/modul", modulController.createModul);
router.put("/modul/:id", modulController.updateModul);
router.delete("/modul/:id", modulController.deleteModul);

module.exports = router;


