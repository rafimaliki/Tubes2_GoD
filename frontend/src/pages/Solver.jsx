import React from "react";
import { useState, useEffect } from "react";
import InputBox from "../components/solver/InputBox";
import SearchBox from "../components/solver/SearchBox";

const Solver = () => {
  const [source, setSource] = useState("");
  const [target, setTarget] = useState("");

  return (
    <div className="bg-black text-white w-full min-h-screen flex items-center justify-center">
      <div className="flex flex-col sm:flex-row items-center ">
        <div className="flex flex-col justify-between h-[18rem]">
          <InputBox val={source} setVal={setSource} label="SOURCE" />
          <InputBox val={target} setVal={setTarget} label="TARGET" />
        </div>
        <div className="flex justify-between h-[18rem] sm:ml-4 mt-5 sm:mt-0">
          <SearchBox source={source} target={target} />
        </div>
      </div>
    </div>
  );
};
export default Solver;
