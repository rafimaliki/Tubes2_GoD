import React from "react";
import { Link, useLocation } from "react-router-dom";

const NavBar = () => {
  const location = useLocation();
  const isActive = (path) => {
    return location.pathname === path;
  };

  return (
    <nav className="border-b border-zinc-700 text-xl bg-black text-white w-full flex items-center justify-between p-4">
      <Link to="/solver" className="font-bold">
        GoD
      </Link>
      <div className="flex font-light text-center">
        <Link
          to="/solver"
          className={`block w-[5rem] hover:font-bold hover:text-white transition duration-200 ease-in-out ${
            isActive("/solver") ? "text-white" : "text-zinc-400"
          }`}
        >
          Solver
        </Link>
        <Link
          to="/about"
          className={`block w-[5rem] hover:font-bold hover:text-white transition duration-200 ease-in-out ${
            isActive("/about") ? "text-white" : "text-zinc-400"
          }`}
        >
          About
        </Link>
      </div>
    </nav>
  );
};

export default NavBar;
