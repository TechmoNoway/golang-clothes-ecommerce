import Header from "@/components/shared/Header";
import { Outlet } from "react-router-dom";

const RootLayout = () => {
  return (
    <>
      <div>
        <div className="w-full h-full">
          <Header />

          <section className="flex flex-1 w-full h-full">
            <Outlet />
          </section>

          {/* <Bottombar /> */}
        </div>
      </div>
    </>
  );
};

export default RootLayout;
