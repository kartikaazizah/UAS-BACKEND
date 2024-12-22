const Modul = require("../models/modul");

exports.getAllModul = async (req, res) => {
  try {
    const modul = await Modul.find();
    res.status(200).json({ data: modul });
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};

exports.createModul = async (req, res) => {
  try {
    const newModul = await Modul.create(req.body);
    res.status(201).json({ message: "Modul created successfully", data: newModul });
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};

exports.updateModul = async (req, res) => {
  try {
    const modul = await Modul.findByIdAndUpdate(req.params.id, req.body, { new: true });
    res.status(200).json({ message: "Modul updated successfully", data: modul });
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};

exports.deleteModul = async (req, res) => {
  try {
    await Modul.findByIdAndDelete(req.params.id);
    res.status(200).json({ message: "Modul deleted successfully" });
  } catch (error) {
    res.status(500).json({ message: error.message });
  }
};


