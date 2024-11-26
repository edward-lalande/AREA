import { Typography, TypographyProps } from '@mui/material';

interface AreaTypographyProps extends TypographyProps {
    text: string;
}

const AreaTypography: React.FC<AreaTypographyProps> = ({ text, ...props }) => {
    return (
        <Typography
            variant="h5"
            fontWeight="bold"
            sx={{
                ...props.sx
            }}
            {...props}
        >
            {text}
        </Typography>
    );
};

export { AreaTypography };
