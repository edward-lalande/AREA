import { Box, BoxProps } from '@mui/material';

const AreaBox: React.FC<BoxProps> = (props) => {
    return (
        <Box
            {...props}
            sx={{
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                justifyContent: "center",
                ...props.sx
            }}
        >
            {props.children}
        </Box>
    );
};
  
export { AreaBox };
