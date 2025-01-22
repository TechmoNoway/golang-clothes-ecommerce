import {
  Heart,
  Search,
  Settings,
  ShoppingBag,
  Users,
} from "lucide-react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";

const Header = () => {
  return (
    <header className="border-b w-full">
      {/* Promo Bar */}
      <div className="bg-gray-100 text-center text-xs font-semibold py-4 flex justify-between items-center px-44">
        <p>Mid-season sale up to 20% OFF. Use code "SALEOFF20"</p>
        <div className="flex space-x-5">
          <div className="text-gray-500">
            Eng <span className="text-xs">&#x25BC;</span>
          </div>
          <div className="text-gray-500">
            USD <span className="text-xs">&#x25BC;</span>
          </div>
        </div>
      </div>

      {/* Main Header */}
      <div className="flex justify-between items-center px-44 py-8">
        {/* Left Section: Search Bar */}
        <div className="flex items-center space-x-2">
          <Input
            type="text"
            placeholder="Search products"
            className="border-none px-3 py-2 w-52 focus:outline-none focus:border-none focus-visible:outline-none  focus-visible:ring-0 focus-visible:border-none font-semibold shadow-none"
          />
          <Button className="p-0 bg-transparent text-black hover:bg-transparent hover:border-transparent">
            <Search className="w-6 h-5" />
          </Button>
        </div>

        {/* Center Section: Logo */}
        <div className="">
          <h1 className="text-2xl font-bold">
            <span className="text-gray-800">Pedona</span>
            <span className="text-red-500">.</span>
          </h1>
        </div>

        {/* Right Section: Icons */}
        <div className="flex items-center space-x-6">
          <div className="flex items-center space-x-4">
            {/* Icons */}
            <Button className="p-0 bg-transparent text-black hover:bg-transparent hover:border-transparent hover:text-red-500">
              <Users className="w-6 h-5" />
            </Button>
            <Button className="p-0 bg-transparent text-black hover:bg-transparent hover:border-transparent hover:text-red-500">
              <Heart className="w-6 h-5" />
            </Button>
            <div className="relative flex">
              <button className="p-0 bg-transparent text-black hover:bg-transparent hover:border-transparent hover:text-red-500">
                <ShoppingBag className="w-6 h-5" />
              </button>
              <span className="absolute -top-2 -right-2 bg-red-500 text-white text-[10px] rounded-full px-2 py-1">
                0
              </span>
            </div>
            <Button className="p-0 bg-transparent text-black hover:bg-transparent hover:border-transparent hover:text-red-500">
              <Settings className="w-6 h-5" />
            </Button>
          </div>
        </div>
      </div>

      {/* Navigation */}
      <nav className="border-t">
        <ul className="flex justify-center space-x-8 text-sm font-medium py-2">
          <li className="text-red-500">HOME</li>
          <li>SHOP</li>
          <li>BLOG</li>
          <li>PORTFOLIO</li>
          <li>PAGES</li>
          <li>CONTACT US</li>
        </ul>
      </nav>
    </header>
  );
};

export default Header;
