const checkRole = (role) => {
  return (req, res, next) => {
    // Simpan role dari user (sesuai dengan implementasi autentikasi Anda)
    const userRole = req.user?.role || "guest"; // Default role jika tidak ada autentikasi

    if (userRole === role) {
      next();
    } else {
      res.status(403).json({ message: "Access denied: Invalid role" });
    }
  };
};

module.exports = checkRole;
