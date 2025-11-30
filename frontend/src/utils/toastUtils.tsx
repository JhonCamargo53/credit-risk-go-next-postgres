import GenericBasicToast from '@/components/common/toast/GenericBasicToast';
import axios from 'axios';
import toast, { ToastPosition } from 'react-hot-toast';
import { MdError, MdCheckCircle, MdWarning, MdInfo } from 'react-icons/md';


export const generateAxiosErrorToast = (error: unknown,
  title: string = "Error",
  message = "Ha ocurrido un error inesperado",
  duration = 3000,
  position: ToastPosition = 'bottom-right'
) => {


  if (axios.isAxiosError(error)) {

    if (error.code === 'ERR_NETWORK') {
      const errorMessage = 'No se ha podido establecer conexión con el servidor.'
      showErrorToast(title, errorMessage, position, duration);
      return errorMessage
    }
    
    if (error.response?.data) {
      message = error.response.data;
    } 

    if(error.response?.data?.error){
        message = error.response.data.error;
    }
    
  } else if (error instanceof Error) {
    title = error.name;
    message = error.message;
  } else {
    message = JSON.stringify(error);
  }

  showErrorToast(title, message, position, duration);

  return message
};

export const showErrorToast = (title: string, message: string, position: ToastPosition = 'bottom-right', duration = 3000) => {
  toast.custom((t) => (
    <GenericBasicToast
      title={title}
      message={message}
      icon={<MdError className="w-5 h-5" />}
      borderColor="border-red-900"
      iconColor="text-red-800"
      textColor="text-red-700"
      shadowColor="red"
      bgColor="bg-white/90"
      onClose={() => toast.dismiss(t.id)}
    />
  ), { position, duration });
};

export const showSuccessToast = (title: string, message: string = '', position: ToastPosition = 'bottom-right', duration = 3000) => {
  toast.custom((t) => (
    <GenericBasicToast
      title={title}
      message={message}
      icon={<MdCheckCircle className="w-5 h-5" />}
      borderColor="border-green-700"
      iconColor="text-green-700"
      textColor="text-green-700"
      shadowColor="green"
      bgColor="bg-white/90"
      onClose={() => toast.dismiss(t.id)}
    />
  ), { position, duration });
};

export const showWarningToast = (title: string, message: string, position: ToastPosition = 'bottom-right', duration = 3000) => {
  toast.custom((t) => (
    <GenericBasicToast
      title={title}
      message={message}
      icon={<MdWarning className="w-5 h-5" />}
      borderColor="border-yellow-600"
      iconColor="text-yellow-600"
      textColor="text-yellow-600"
      shadowColor="yellow"
      bgColor="bg-white/90"
      onClose={() => toast.dismiss(t.id)}
    />
  ), { position, duration });
};

export const showInfoToast = (title: string, message: string, position: ToastPosition = 'bottom-right', duration = 3000) => {
  toast.custom((t) => (
    <GenericBasicToast
      title={title}
      message={message}
      icon={<MdInfo className="w-5 h-5" />}
      borderColor="border-blue-600"
      iconColor="text-blue-600"
      textColor="text-blue-600"
      shadowColor="blue"
      bgColor="bg-white/90"
      onClose={() => toast.dismiss(t.id)}
    />
  ), { position, duration });
};
