import React from "react";
import { Link, useLocation } from "react-router-dom";

const NavBar = () => {
  const location = useLocation();
  const isActive = (path) => {
    return location.pathname === path || location.pathname.startsWith(path);
  };
  return (
    <nav className="border-b border-zinc-700 text-xl bg-black text-white w-full flex items-center justify-between p-4">
      <Link to="/" className="font-bold">
        GoD
      </Link>
      <div className="font-light">
        <Link
          to="/game"
          className={`mr-4 ${
            isActive("/game") ? "text-white" : "text-zinc-400"
          }`}
        >
          Game
        </Link>
        <Link
          to="/about"
          className={`${isActive("/about") ? "text-white" : "text-zinc-400"}`}
        >
          About
        </Link>
      </div>
    </nav>
  );
};

export default NavBar;
