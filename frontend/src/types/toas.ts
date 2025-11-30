export interface GenericToast {
    title: string
    icon?: React.ReactNode,
    message: string,
    bgColor?: string,
    textColor?: string
    borderColor?: string
    iconColor?: string
    shadowColor?: 'green' | 'red' | 'yellow' | 'blue',
    onClose?: () => void
}