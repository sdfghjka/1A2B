import React from "react";
import Header from "./header";
function Layout({ children }) {
  return (
    <>
      <Header />
      <main style={{ paddingTop: "80px" }}>{children}</main>
    </>
  );
}

export default Layout;



