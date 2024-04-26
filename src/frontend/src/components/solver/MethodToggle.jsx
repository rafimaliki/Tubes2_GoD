import React from "react";

const MethodToggle = ({ label, method, setMethod }) => {
  return (
    <div className="flex items-center justify-center  w-[5rem] mx-3">
      <button
        className={`mr-3 w-3 h-3 rounded-full  flex justify-center items-center border-white border-2 ${
          method == label ? "bg-black" : "bg-white"
        }
        }`}
        onClick={() => setMethod(label)}
      ></button>
      <p className="text-center font-bold text-xl">{label}</p>
    </div>
  );
};

export default MethodToggle;
