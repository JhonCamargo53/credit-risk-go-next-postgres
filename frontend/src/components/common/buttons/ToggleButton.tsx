import React from 'react'
import { HashLoader } from 'react-spinners';
interface ToggleButtonProps {
    checked: boolean
    handleClick: () => void
    loading?: boolean
    loadingComponent?: React.ReactNode;
}

const ToggleButton: React.FC<ToggleButtonProps> = ({ checked, handleClick, loading = false, loadingComponent = <HashLoader color="var(--brand-primary)" size={25} />,
}) => {
    return (
        <>
            {loading ? (
                <div className="flex items-center justify-center">
                    {loadingComponent}
                </div>
            ) : (
                <div className="flex items-center justify-center content-center">
                    <label className="cursor-pointer"                    >
                        <input
                            type="checkbox"
                            className="sr-only peer"
                            checked={checked}
                            onClick={() => handleClick()}
                            readOnly
                        />
                        <div
                            className="relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4
        peer-focus:ring-primary/30 dark:bg-neutral-dark 
        peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full
        peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:start-[2px]
        after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 
        after:transition-all dark:border-gray-600 peer-checked:bg-red-600 dark:peer-checked:bg-primary"
                        ></div>
                    </label>
                </div>
            )}
        </>
    );

}

export default ToggleButton