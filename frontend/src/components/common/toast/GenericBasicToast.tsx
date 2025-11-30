import { GenericToast } from '@/types/toas';
import React from 'react';
import { CgClose } from "react-icons/cg";
import { MdCheckCircle } from 'react-icons/md';
import { twMerge } from 'tailwind-merge';

const getShadowClass = (shadowColor: string) => {
  switch (shadowColor) {
    case 'green':
      return 'shadow-[0_0_10px_rgba(22,163,74,0.3)]';
    case 'red':
      return 'shadow-[0_0_10px_rgba(239,68,68,0.3)]';
    case 'yellow':
      return 'shadow-[0_0_10px_rgba(202,138,4,0.3)]';
    case 'blue':
      return 'shadow-[0_0_10px_rgba(37,99,235,0.3)]';
    default:
      return 'shadow-[0_0_10px_rgba(22,163,74,0.3)]';
  }
};

const getBgColorFromBorder = (borderColor: string) => {
  switch (borderColor) {
    case 'border-green-700':
      return 'bg-green-700';
    case 'border-red-900':
      return 'bg-red-900';
    case 'border-yellow-600':
      return 'bg-yellow-600';
    case 'border-blue-600':
      return 'bg-blue-600';
    default:
      return 'bg-green-700';
  }
};

const GenericBasicToast: React.FC<GenericToast> = ({
  title,
  message,
  icon = <MdCheckCircle />,
  iconColor = 'text-green-700',
  borderColor = 'border-green-700',
  textColor = 'text-green-700',
  bgColor = 'bg-white/90',
  shadowColor = 'green', // default to green
  onClose = () => { }
}) => {

  const messageFormat = (message: string) => {
    return message.endsWith('.') ? message : message + "."
  }
  return (
    <div
      className={twMerge(
        'rounded-md overflow-hidden',
        getShadowClass(shadowColor),
        bgColor,
        borderColor
      )}
    >
      <div className="flex items-start justify-between p-3">
        <div className="flex items-start gap-2 w-full">
          <div className={iconColor}>
            {icon}
          </div>
          <div className="flex flex-col">
            <span className={twMerge('font-bold text-sm', textColor)}>{title}</span>
            {message && <div>
              {message.split('/*./split/.*/').map((messages, index) => (
                <div key={index} className="text-sm text-gray-600 break-words">
                  - {messageFormat(messages)}
                </div>
              ))}
            </div>}

          </div>
        </div>
        <button onClick={() => onClose()} className="ml-3 text-gray-400 hover:text-gray-600">
          <CgClose />
        </button>
      </div>
      <div className={twMerge('h-1 w-full', getBgColorFromBorder(borderColor))}></div>
    </div>
  );
};

export default GenericBasicToast;
