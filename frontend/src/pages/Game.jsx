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
        <div className="flex flex-col mt-5">
          <SearchButton
            source={source}
            target={target}
            endpoint="search_bfs"
            label="Breadth First Search"
          />
          <SearchButton
            source={source}
            target={target}
            endpoint="search_ids"
            label="Iterative Deepening Search"
          />
        </div>
      </div>
    </div>
  );
};
export default Game;
