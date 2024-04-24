import React from "react";

const Suggestion = ({ title, setVal }) => {
  const handleClick = () => {
    setVal(title);
    console.log(title);
  };
  return (
    <div className="border border-white rounded-sm w-[17rem]">
      <button
        className="h-full w-full p-0 pl-1 text-gray-300 bg-zinc-900 hover:bg-zinc-800"
        onClick={handleClick}
      >
        {title}
      </button>
    </div>
  );
};

export default Suggestion;
