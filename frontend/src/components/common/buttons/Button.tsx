import React from 'react';
import { HashLoader } from 'react-spinners';

interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  loading?: boolean;
  children: React.ReactNode;
  variant?: 'primary' | 'secondary';
  loadingComponent?: React.ReactNode;
  iconLeft?: React.ReactNode;
  iconRight?: React.ReactNode;
}

const Button: React.FC<ButtonProps> = ({
  loading = false,
  children,
  variant = 'primary',
  disabled,
  loadingComponent = <HashLoader color="white" size={25} />,
  iconLeft,
  iconRight,
  ...props
}) => {

  const getVariantClass = () => {
    switch (variant) {
      case 'secondary':
        return 'btn-secondary'
      default:
        return 'btn-primary';
    }
  };

  return (
    <button
      disabled={loading || disabled}
      {...props}
      className={`${getVariantClass()} flex items-center justify-center gap-2 ${props.className || ''}`}
    >
      {loading ? (
        <div className="flex justify-center items-center">{loadingComponent}</div>
      ) : (
        <>
          {iconLeft && <span className="flex items-center">{iconLeft}</span>}
          <span>{children}</span>
          {iconRight && <span className="flex items-center">{iconRight}</span>}
        </>
      )}
    </button>
  );
};

export default Button;
