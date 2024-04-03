import React from "react";
import { useState } from "react";
import InputBox from "../components/InputBox";
import SearchButton from "../components/SearchButton";

const Game = () => {
  const [source, setSource] = useState("");
  const [target, setTarget] = useState("");
  return (
    <div className="bg-black text-white w-full min-h-screen flex items-center justify-center">
      <div className="flex flex-col items-center">
        <div className="flex">
          <InputBox val={source} setVal={setSource} label="Source" />
          <InputBox val={target} setVal={setTarget} label="Target" />
        </div>
        <SearchButton source={source} target={target} />
      </div>
    </div>
  );
};
export default Game;
