import { Typography, TypographyProps } from '@mui/material';

interface AreaTypographyProps extends TypographyProps {
    text: string;
}

const AreaTypography: React.FC<AreaTypographyProps> = ({ text, ...props }) => {
    return (
        <Typography
            variant="h5"
            fontWeight="bold"
            {...props}
            sx={{
                ...props.sx
            }}
        >
            {text}
        </Typography>
    );
};

const AreasTypography: React.FC<AreaTypographyProps> = ({ text, ...props }) => {
    return (
        <Typography
            variant="h6"
            {...props}
            sx={{
                ...props.sx
            }}
        >
            {text}
        </Typography>
    );
};

export { AreaTypography, AreasTypography };
