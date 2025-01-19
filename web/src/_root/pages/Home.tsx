import React from "react";

const HeroSection = () => {
  return (
    <section className="bg-black text-white relative w-full h-full">
      {/* Hero Content */}
      <div className="max-w-7xl mx-auto px-4 py-20 text-center">
        <div className="inline-block bg-red-600 text-xs text-white uppercase px-3 py-1 rounded">
          Hot Trending
        </div>
        <h1 className="text-4xl font-bold mt-4">
          Sale Offer 20% Off This Week
        </h1>
        <p className="text-5xl font-extrabold mt-2">
          Travel Handbags
        </p>
        <button className="mt-6 px-6 py-3 bg-white text-black font-bold rounded hover:bg-gray-200">
          Shopping Now
        </button>
      </div>

      {/* Background Image */}
      <div className="absolute inset-0 opacity-40">
        <img
          src="https://i.ibb.co/QfHYCm5/banner.jpg"
          alt="Background"
          className="w-full h-full object-cover"
        />
      </div>

      {/* Navigation Arrows */}
      <div className="absolute inset-y-0 left-0 flex items-center">
        <button className="text-white text-2xl p-3 bg-black bg-opacity-50 rounded-full">
          ❮
        </button>
      </div>
      <div className="absolute inset-y-0 right-0 flex items-center">
        <button className="text-white text-2xl p-3 bg-black bg-opacity-50 rounded-full">
          ❯
        </button>
      </div>
    </section>
  );
};

export default HeroSection;
