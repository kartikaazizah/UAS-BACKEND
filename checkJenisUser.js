const User = require("../models/user");

const checkJenisUser = (jenisUser) => {
  return async (req, res, next) => {
    try {
      const user = await User.findById(req.user.id);
      if (user && user.jenis_user === jenisUser) {
        next();
      } else {
        return res.status(403).json({ message: "Access Denied: Invalid Jenis User" });
      }
    } catch (error) {
      return res.status(500).json({ message: error.message });
    }
  };
};

module.exports = checkJenisUser;
