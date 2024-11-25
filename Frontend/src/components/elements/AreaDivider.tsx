import { Box, Typography, Divider, DividerProps } from '@mui/material';

interface AreaTextDividerProps extends DividerProps {
    text: string;
}

const AreaTextDivider: React.FC<AreaTextDividerProps> = ({ text, ...props }) => {
    return (
        <Box
            sx={{
                display: "flex",
                alignItems: "center",
                width: "100%",
                maxWidth: 400,
                my: 3,
                ...props.sx
            }}
            {...props}
        >
            <Divider sx={{ flex: 1}} />
            <Typography sx={{ px: 2, fontSize: "0.9rem", color: "#888" }}>{text}</Typography>
            <Divider sx={{ flex: 1}} />
        </Box>
    );
};

const AreaDivider: React.FC<DividerProps> = (props) => {
    return (
        <Box
            sx={{
                display: "flex",
                alignItems: "center",
                width: "100%",
                ...props.sx
            }}
            {...props}
        >
            <Divider sx={{ flex: 1 }} />
        </Box>
    );
};
  
export { AreaDivider, AreaTextDivider };
