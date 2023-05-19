import React from "react";
import { Link } from "react-router-dom";

export default function HomeSiderLogo() {
  return (
    <Link to={"/"}>
      <div
        style={{
          width: "100%",
          background: "#fff",
          height: "60px",
          display: "flex",
          alignItems: "center",
        }}
      >
        <img
          style={{
            width: "100%",
            background: "#fff",
            padding: "1rem 0",
          }}
          alt="Zalora logo"
          src="https://upload.wikimedia.org/wikipedia/commons/thumb/2/23/Zalora_Group_logo.svg/2560px-Zalora_Group_logo.svg.png"
        />
      </div>
    </Link>
  );
}
