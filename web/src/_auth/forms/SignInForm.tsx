// import { z } from "zod";

// import { FaGithub } from "react-icons/fa6";
// import { FcGoogle } from "react-icons/fc";
// import { zodResolver } from "@hookform/resolvers/zod";
// import { useForm } from "react-hook-form";
// import { Link } from "react-router-dom";
// import {
//   Form,
//   FormControl,
//   FormField,
//   FormItem,
//   FormLabel,
//   FormMessage,
// } from "@/components/ui/form";
// import { Input } from "@/components/ui/input";
// import PasswordInput from "@/components/shared/PasswordInput";
// import { Button } from "@/components/ui/button";

// const formSchema = z.object({
//   username: z.string().min(4, {
//     message: "Username must be at least 4 characters.",
//   }),
//   password: z.string().min(1, {
//     message: "Dont let the password empty.",
//   }),
// });

// const SignInForm = () => {
//   const form = useForm<z.infer<typeof formSchema>>({
//     resolver: zodResolver(formSchema),
//     defaultValues: {
//       username: "",
//       password: "",
//     },
//   });

//   async function onSubmit(values: z.infer<typeof formSchema>) {
//     console.log(values);
//   }

//   return (
//     <>
//       <Form {...form}>
//         <div className="sm:w-420 flex flex-col">
//           <h2 className="text-[30px] text-blue-900 font-bold">
//             Sign In
//           </h2>

//           <form
//             onSubmit={form.handleSubmit(onSubmit)}
//             className="flex flex-col gap-5 w-full mt-12"
//           >
//             <FormField
//               control={form.control}
//               name="username"
//               render={({ field }) => (
//                 <FormItem>
//                   <FormLabel className="text-blue-900">
//                     Username
//                   </FormLabel>
//                   <FormControl>
//                     <Input
//                       placeholder="Username"
//                       className="py-7 w-96 font-medium text-base bg-gray-100 border-none"
//                       {...field}
//                     />
//                   </FormControl>
//                   <FormMessage />
//                 </FormItem>
//               )}
//             />

//             <FormField
//               control={form.control}
//               name="password"
//               render={({ field }) => (
//                 <FormItem>
//                   <FormLabel className="text-blue-900">
//                     Password
//                   </FormLabel>
//                   <FormControl>
//                     <PasswordInput
//                       placeholder="Password"
//                       className="py-7 w-96 font-medium text-base bg-gray-100 border-none"
//                       {...field}
//                     />
//                   </FormControl>
//                   <FormMessage />
//                 </FormItem>
//               )}
//             />

//             <Button
//               type="submit"
//               className="bg-blue-500 py-6 text-xl"
//             >
//               Signin
//             </Button>

//             <hr />

//             <p className="flex justify-center">Or sign in with</p>

//             <div className="flex justify-center space-x-5">
//               <Button
//                 className="bg-white w-full text-black border-[1px] border-black space-x-2 hover:text-white"
//                 type="button"
//               >
//                 <FaGithub className="w-6 h-6" /> <p>Github</p>
//               </Button>
//               <Button
//                 className="bg-white w-full text-black border-[1px] border-black space-x-2 hover:text-white"
//                 type="button"
//               >
//                 <FcGoogle className="w-6 h-6" /> <p>Google</p>
//               </Button>
//             </div>
//             <hr />

//             <div className="flex justify-between">
//               <p className="font-semibold text-slate-600">
//                 Don't have an account?
//               </p>
//               <Link
//                 to={"/sign-up"}
//                 className="font-semibold text-blue-600 hover:text-black"
//               >
//                 Sign up for free
//               </Link>
//             </div>
//           </form>
//         </div>
//       </Form>
//     </>
//   );
// };

// export default SignInForm;
import React from "react";

const LoginPage: React.FC = () => {
  return (
    <div className="flex h-screen bg-white">
      {/* Left Section */}
      <div className="flex flex-col justify-between w-full max-w-md px-8 py-12 mx-auto">
        {/* Logo */}
        <div className="text-center">
          <h1 className="text-3xl font-bold text-black">somobile</h1>
        </div>

        {/* Login Form */}
        <div>
          <h2 className="text-lg font-semibold text-center text-gray-800">
            Login
          </h2>
          <form className="flex flex-col mt-6 space-y-4">
            <input
              type="email"
              placeholder="Email address"
              className="w-full px-4 py-2 text-sm border rounded-full focus:ring-2 focus:ring-blue-500 focus:outline-none"
            />
            <input
              type="password"
              placeholder="Password"
              className="w-full px-4 py-2 text-sm border rounded-full focus:ring-2 focus:ring-blue-500 focus:outline-none"
            />
            <a
              href="#"
              className="text-sm text-gray-500 hover:underline text-center"
            >
              Forgot my password
            </a>
            <button
              type="submit"
              className="w-full py-2 text-sm font-semibold text-white bg-black rounded-full hover:bg-gray-800"
            >
              Continue to somobile
            </button>
          </form>
        </div>

        {/* Social Media Icons */}
        <div className="flex items-center justify-center space-x-4">
          <a href="#" className="text-gray-600 hover:text-black">
            <i className="fab fa-facebook-f"></i>
          </a>
          <a href="#" className="text-gray-600 hover:text-black">
            <i className="fab fa-twitter"></i>
          </a>
          <a href="#" className="text-gray-600 hover:text-black">
            <i className="fab fa-linkedin-in"></i>
          </a>
        </div>

        {/* Footer */}
        <footer className="text-xs text-center text-gray-400">
          <p>© 2025 somobile · Terms · Legal</p>
        </footer>
      </div>

      {/* Right Section */}
      <div className="hidden md:block w-1/2 bg-gray-100 relative">
        <div className="absolute inset-0 flex items-center justify-center">
          <div className="w-96 h-96 overflow-hidden rounded-full">
            <img
              src="/path-to-your-image.jpg"
              alt="Yellow van by the sea"
              className="object-cover w-full h-full"
            />
          </div>
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
