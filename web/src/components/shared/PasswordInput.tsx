import React, { forwardRef, useState } from "react";
import { AiOutlineEyeInvisible, AiOutlineEye } from "react-icons/ai";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";

const PasswordInput = forwardRef<
  HTMLInputElement,
  React.InputHTMLAttributes<HTMLInputElement>
>(({ className, ...props }, ref) => {
  const [showPassword, setShowPassword] = useState(false);
  const disabled =
    props.value === "" || props.value === undefined || props.disabled;

  return (
    <div className="relative">
      <Input
        type={showPassword ? "text" : "password"}
        className={cn("pr-10", className)}
        ref={ref}
        {...props}
      />
      <Button
        type="button"
        variant="ghost"
        size="sm"
        className="absolute right-0 top-0 w-14 h-full px-3 py-2  hover:bg-transparent cursor-pointer bg-transparent flex hover:border-none focus:outline-none"
        onClick={() => setShowPassword(!showPassword)}
        disabled={disabled}
      >
        {showPassword && !disabled ? (
          <AiOutlineEye
            className="text-2xl items-center justify-center"
            aria-hidden="true"
          />
        ) : (
          <AiOutlineEyeInvisible
            className="text-2xl items-center justify-center"
            aria-hidden="true"
          />
        )}
      </Button>
    </div>
  );
});
PasswordInput.displayName = "PasswordInput";

export default PasswordInput;
