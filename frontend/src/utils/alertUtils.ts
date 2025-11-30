import Swal, { SweetAlertIcon } from 'sweetalert2'

export const confirmActionAlert = async (
    title: string,
    message: string,
    icon: SweetAlertIcon,
    confirmButtonText: string = "Confirmar",
    cancelButtonText: string = "Cancelar"
): Promise<boolean> => {

    const formattedMessage = message.replace(/"([^"]+)"/g, '<strong>$1</strong>');

    return new Promise<boolean>((resolve) => {
        Swal.fire({
            title: title,
            html: formattedMessage,
            icon: icon,
            showCancelButton: true,
            confirmButtonColor: "var(--brand-primary)",
            cancelButtonColor: "var(--brand-neutral-dark)",
            confirmButtonText: confirmButtonText,
            cancelButtonText: cancelButtonText,
        }).then((result) => {
            if (result.isConfirmed) {
                resolve(true);
            } else {
                resolve(false);
            }
        });
    });
};