import React from "react";

const Header = () => {
  return (
    <header className="bg-gray-900 text-white">
      {/* Top Bar */}
      <div className="flex justify-between items-center px-4 py-2 text-sm border-b border-gray-700">
        <p>
          Mid-season sale up to <strong>20% OFF</strong>. Use code
          "SALEOFF20"
        </p>
        <div className="flex space-x-4">
          <span>Eng</span>
          <span>USD</span>
        </div>
      </div>

      {/* Navigation */}
      <div className="flex justify-between items-center px-4 py-4">
        {/* Logo */}
        <div className="text-xl font-bold">Pedona</div>

        {/* Navigation Links */}
        <nav className="hidden md:flex space-x-6 text-sm">
          <a href="#" className="hover:text-red-500">
            Home
          </a>
          <a href="#" className="hover:text-red-500">
            Shop
          </a>
          <a href="#" className="hover:text-red-500">
            Blog
          </a>
          <a href="#" className="hover:text-red-500">
            Portfolio
          </a>
          <a href="#" className="hover:text-red-500">
            Pages
          </a>
          <a href="#" className="hover:text-red-500">
            Contact Us
          </a>
        </nav>

        {/* Icons */}
        <div className="flex space-x-4 text-lg">
          <i className="fas fa-user"></i>
          <i className="fas fa-heart"></i>
          <i className="fas fa-shopping-cart relative">
            <span className="absolute -top-2 -right-2 bg-red-500 text-white text-xs w-5 h-5 rounded-full flex items-center justify-center">
              0
            </span>
          </i>
          <i className="fas fa-cog"></i>
        </div>
      </div>
    </header>
  );
};

export default Header;
