import { z } from "zod";

import { FaGithub } from "react-icons/fa6";
import { FcGoogle } from "react-icons/fc";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { Link } from "react-router-dom";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import PasswordInput from "@/components/shared/PasswordInput";
import { Button } from "@/components/ui/button";

const formSchema = z.object({
  username: z.string().min(4, {
    message: "Username must be at least 4 characters.",
  }),
  password: z.string().min(1, {
    message: "Dont let the password empty.",
  }),
});

const SignInForm = () => {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    console.log(values);
  }

  return (
    <>
      <Form {...form}>
        <div className="sm:w-420 flex flex-col">
          <h2 className="text-[30px] text-blue-900 font-bold">
            Sign In
          </h2>

          <form
            onSubmit={form.handleSubmit(onSubmit)}
            className="flex flex-col gap-5 w-full mt-12"
          >
            <FormField
              control={form.control}
              name="username"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-blue-900">
                    Username
                  </FormLabel>
                  <FormControl>
                    <Input
                      placeholder="Username"
                      className="py-7 w-96 font-medium text-base bg-gray-100 border-none"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="password"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="text-blue-900">
                    Password
                  </FormLabel>
                  <FormControl>
                    <PasswordInput
                      placeholder="Password"
                      className="py-7 w-96 font-medium text-base bg-gray-100 border-none"
                      {...field}
                    />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />

            <Button
              type="submit"
              className="bg-blue-500 py-6 text-xl"
            >
              Signin
            </Button>

            <hr />

            <p className="flex justify-center">Or sign in with</p>

            <div className="flex justify-center space-x-5">
              <Button
                className="bg-white w-full text-black border-[1px] border-black space-x-2 hover:text-white"
                type="button"
              >
                <FaGithub className="w-6 h-6" /> <p>Github</p>
              </Button>
              <Button
                className="bg-white w-full text-black border-[1px] border-black space-x-2 hover:text-white"
                type="button"
              >
                <FcGoogle className="w-6 h-6" /> <p>Google</p>
              </Button>
            </div>
            <hr />

            <div className="flex justify-between">
              <p className="font-semibold text-slate-600">
                Don't have an account?
              </p>
              <Link
                to={"/sign-up"}
                className="font-semibold text-blue-600 hover:text-black"
              >
                Sign up for free
              </Link>
            </div>
          </form>
        </div>
      </Form>
    </>
  );
};

export default SignInForm;
