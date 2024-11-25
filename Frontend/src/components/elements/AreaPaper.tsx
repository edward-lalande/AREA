import { Paper, PaperProps } from '@mui/material';

const AreaPaper: React.FC<PaperProps> = (props) => {
    return (
        <Paper
            elevation={6}
            {...props}
            sx={{
                ...props.sx,
                display: "flex",
                flexDirection: "column",
                alignItems: "center",
                borderRadius: 5,
                backgroundColor: "#FFF",
                boxShadow: "0 4px 12px rgba(0, 0, 0, 0.1)",
            }}
		>
            {props.children}
        </Paper>
    );
};
  
export { AreaPaper };
