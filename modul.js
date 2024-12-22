const mongoose = require("mongoose");

const ModulSchema = new mongoose.Schema({
  nm_modul: { type: String, required: true },
  kategori: { type: mongoose.Schema.Types.ObjectId, ref: "Kategori" },
  is_aktif: { type: Boolean, default: true },
  gbr_icon: { type: String },
});

module.exports = mongoose.model("Modul", ModulSchema);
