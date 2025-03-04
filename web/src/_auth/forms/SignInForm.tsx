import { Button } from "@/components/ui/button";
import { Form, FormControl, FormDescription, FormField, FormItem, FormLabel, FormMessage } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

const formSchema = z.object({
  username: z.string().min(2).max(50),
});
const SignInForm = () => {
  // 1. Define your form.
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
    },
  });

  // 2. Define a submit handler.
  function onSubmit(values: z.infer<typeof formSchema>) {
    // Do something with the form values.
    // ✅ This will be type-safe and validated.
    console.log(values);
  }

  return (
    <div className="flex h-screen rounded-lg shadow-lg">
      <div className="flex flex-col justify-center items-center w-[400px] md:h-full lg:h-full md:w-1/2 lg:w-1/2 p-10 bg-white">
        <div className="w-full max-w-sm">
          <h1 className="text-2xl font-bold text-orange-500 mb-6">
            Golend
          </h1>

          <p className="text-gray-500 mb-2">Welcome back !!!</p>
          <h2 className="text-3xl font-bold mb-6">Sign in</h2>

          <div className="mb-4">
            <label
              className="block text-gray-500 text-sm mb-2"
              htmlFor="email"
            >
              Email
            </label>
            <input
              type="email"
              id="email"
              placeholder="abc@gmail.com"
              className="w-full px-4 py-3 border border-gray-300 rounded-lg bg-[#FFF6F3] focus:outline-none"
            />
          </div>

          <div className="mb-4">
            <div className="flex justify-between">
              <label
                className="block text-gray-500 text-sm mb-2"
                htmlFor="password"
              >
                Password
              </label>
              <a
                href="#"
                className="text-sm text-orange-500 hover:underline ml-2"
              >
                Forgot Password?
              </a>
            </div>
            <div className="flex justify-between">
              <input
                type="password"
                id="password"
                placeholder="********"
                className="w-full px-4 py-3 border border-gray-300 rounded-lg bg-[#FFF6F3] focus:outline-none"
              />
            </div>
          </div>

          {/* Sign-In Button */}
          <Button className="w-full py-3 bg-orange-500 text-white rounded-lg shadow-md hover:bg-orange-600">
            SIGN IN →
          </Button>

          {/* Sign-Up Link */}
          <p className="mt-4 text-sm text-gray-500 text-center">
            I don't have an account?
            <a href="#" className="text-orange-500 hover:underline">
              Sign up
            </a>
          </p>
        </div>
      </div>

      <div className="lg:flex md:flex md:w-1/2 lg:w-1/2 justify-center items-center hidden bg-[#FFEFE8]">
        <div className="relative">
          <img
            src="https://www.atnbazaar.com/images/login.png"
            alt="Shopping Illustration"
            className=""
          />
        </div>
      </div>
      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="space-y-8"
        >
          <FormField
            control={form.control}
            name="username"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Username</FormLabel>
                <FormControl>
                  <Input placeholder="shadcn" {...field} />
                </FormControl>
                <FormDescription>
                  This is your public display name.
                </FormDescription>
                <FormMessage />
              </FormItem>
            )}
          />
          <Button type="submit">Submit</Button>
        </form>
      </Form>
    </div>
  );
};

export default SignInForm;
