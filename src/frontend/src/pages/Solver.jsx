import React from "react";
import { useState, useEffect } from "react";
import InputBox from "../components/solver/InputBox";
import SearchBox from "../components/solver/SearchBox";
import ResultContainer from "../components/solver/ResultContainer";

const Solver = () => {
  const [source, setSource] = useState("");
  const [target, setTarget] = useState("");
  const [result, setResult] = useState({
    path: [],
    duration: 0,
    searched: 0,
  });

  return (
    <div
      className="bg-black text-white w-full flex flex-col items-center justify-center"
      style={{ minHeight: "calc(100vh - 4rem)", height: "100%" }}
    >
      <div className="my-4">
        <div className="flex flex-col sm:flex-row items-center ">
          <div className="flex flex-col justify-between h-[18rem]">
            <InputBox val={source} setVal={setSource} label="SOURCE" />
            <InputBox val={target} setVal={setTarget} label="TARGET" />
          </div>
          <div className="flex justify-between h-[18rem] sm:ml-4 mt-5 sm:mt-0">
            <SearchBox source={source} target={target} setResult={setResult} />
          </div>
        </div>
        <div className="mt-4">
          {result.checked == -1 && (
            <p className="text-center text-red-400">Invalid Source Wiki</p>
          )}
          {result.checked == -2 && (
            <p className="text-center text-red-400">Invalid Target Wiki</p>
          )}
          {result.checked == -3 && (
            <p className="text-center text-red-400">Fail Scrapping Wiki</p>
          )}
          {result.path.length != 0 && <ResultContainer result={result} />}
        </div>
      </div>
    </div>
  );
};
export default Solver;
