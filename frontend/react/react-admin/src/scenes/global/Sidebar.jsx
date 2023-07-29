import { Box, useTheme } from "@mui/material";
import { useState } from "react";
import { Sidebar } from "react-pro-sidebar";

const Sidebar = () => {
	const theme = useTheme();
	const colors = tokens(theme.palette.mode);
	const [isCollapsed, setIsCollapsed] = useState(false);
	const [selected, setSelected] = useState("Dashboard");

	return <Box></Box>;
};

export default Sidebar;
