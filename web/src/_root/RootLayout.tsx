import { Outlet } from "react-router-dom";

const RootLayout = () => {
  return (
    <>
      <div>
        <div className="w-full md:flex">
          {/* <Topbar />
          <LeftSidebar /> */}

          <section className="flex flex-1 h-full">
            <Outlet />
          </section>

          {/* <Bottombar /> */}
        </div>
      </div>
    </>
  );
};

export default RootLayout;
