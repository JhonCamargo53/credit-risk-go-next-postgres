import React from 'react';
import { HashLoader } from 'react-spinners'
interface IconButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
    loading?: boolean;
    children: React.ReactNode;
    loadingComponent?: React.ReactNode
}

const IconButton: React.FC<IconButtonProps> = ({ loading = false, children, disabled, loadingComponent = <HashLoader color='white' size={25} />, ...props }) => {
    return (
        <button
            disabled={loading || disabled}
            {...props}
            className={`btn-primary ${props.className || ''}`}
        >
            {loading ? <div className='flex justify-center'>{loadingComponent}</div> : children}
        </button>
    );
};

export default IconButton;
