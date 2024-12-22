const express = require("express");
const mongoose = require("mongoose");
const bodyParser = require("body-parser");
const adminRoutes = require("./routes/admin");

const app = express();
const PORT = 5000;

// Middleware
app.use(bodyParser.json());

// Routes
app.use("/admin", adminRoutes);

// Connect to MongoDB
mongoose
  .connect("mongodb://localhost:27017/SSO_UAS_BEKEN", { useNewUrlParser: true, useUnifiedTopology: true })
  .then(() => console.log("MongoDB connected"))
  .catch((err) => console.error("MongoDB connection failed:", err));

app.listen(PORT, () => console.log(`Server running on http://localhost:${PORT}`));
