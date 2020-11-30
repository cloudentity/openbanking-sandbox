import React from 'react';
import {makeStyles} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import {Typography} from "@material-ui/core";

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
  withHover: {
    '&:hover': {
      cursor: 'pointer'
    }
  }
});

function createData(name: string, id: string) {
  return {name, id};
}

const rows = [
  createData('Account 123', '123'),
  createData('Account 456', '456'),
];

export default ({onRowClick}) => {
  const classes = useStyles();

  return (
    <>
      <Typography variant={'h4'}>Accounts list</Typography>
      <TableContainer component={Paper} style={{marginTop: 32}}>
        <Table className={classes.table} aria-label="Accounts table">
          <TableHead>
            <TableRow>
              <TableCell>Name</TableCell>
              <TableCell align="right">ID</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {rows.map((row) => (
              <TableRow key={row.id} hover className={classes.withHover} onClick={onRowClick(row)}>
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="right">{row.id}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}
