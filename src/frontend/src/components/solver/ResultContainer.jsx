import React from "react";
import Result from "./Result";

const ResultContainer = ({ result }) => {
  const renderResult = () => {
    if (result.path != null) {
      return (
        <>
          <p className="font-bold text-xl">RESULT</p>
          {result.path.map((wiki, idx) => (
            <Result key={idx} id={idx + 1} title={wiki.Title} url={wiki.URL} />
          ))}
        </>
      );
    }
  };
  return (
    <div className="flex flex-col items-center justify-center sm:w-[36rem] w-[20rem]">
      {renderResult()}
    </div>
  );
};

export default ResultContainer;
