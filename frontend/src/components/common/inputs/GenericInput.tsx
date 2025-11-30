import React, { HTMLInputTypeAttribute, useState } from 'react'
import { FieldError, UseFormRegisterReturn } from 'react-hook-form'
import { FaEye, FaEyeSlash } from 'react-icons/fa'
import clsx from 'clsx';

interface GenericInputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label?: string
  placeholder: string
  type: HTMLInputTypeAttribute
  error?: FieldError
  register: UseFormRegisterReturn
  variant?: 'primary' | 'secondary'
  iconLeft?: React.ReactNode
  iconRight?: React.ReactNode
}

const GenericInput: React.FC<GenericInputProps> = ({
  label = '',
  placeholder,
  type,
  error,
  register,
  variant = 'primary',
  iconLeft,
  iconRight,
  ...props
}) => {
  const [showPassword, setShowPassword] = useState(false);
  const isTypePassword = type === 'password';
  const isSecondary = variant === 'secondary';

  return (
    <div className="w-full">
      {label && <label className='text-neutral-dark font-bold mb-1 block'>{label}</label>}

      <div className="relative flex items-center">
        {iconLeft && (
          <div className="absolute left-3 text-neutral-dark">
            {iconLeft}
          </div>
        )}

        <input
          {...props}
          {...register}
          type={showPassword ? 'text' : type}
          placeholder={placeholder}
          className={clsx(
            ' duration-300 w-full',
            variant === 'primary' && [
              'bg-transparent border-b-2 border-neutral-dark text-neutral-dark px-3 py-2 placeholder-neutral-dark/70',
              'focus:bg-white/30 focus:outline-none rounded-t-sm',
              isTypePassword && 'pr-10',
              iconLeft && 'pl-10',
              iconRight && 'pr-10'
            ],
            variant === 'secondary' && [
              'bg-white border border-neutral-dark text-neutral-dark px-3 py-2 rounded-md placeholder-neutral-dark/70',
              'focus:ring-1 focus:outline-none',
              isTypePassword && 'pr-10',
              iconLeft && 'pl-10',
              iconRight && 'pr-10'
            ],
            props.className
          )}
        />

        {/* Password toggle icon */}
        {isTypePassword && (
          <button
            type="button"
            className="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-dark cursor-pointer"
            onClick={() => setShowPassword(!showPassword)}
          >
            {showPassword ? <FaEyeSlash /> : <FaEye />}
          </button>
        )}

        {/* Custom right icon (if not password) */}
        {!isTypePassword && iconRight && (
          <div className="absolute right-3 top-1/2 -translate-y-1/2 text-neutral-dark">
            {iconRight}
          </div>
        )}
      </div>

      {error && (
        <p className="text-error text-sm mt-1 font-bold">{error.message}</p>
      )}
    </div>
  )
}

export default GenericInput;
