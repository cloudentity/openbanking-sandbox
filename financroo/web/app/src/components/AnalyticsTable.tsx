import React, {useEffect} from 'react';
import clsx from 'clsx';
import {createStyles, lighten, makeStyles, Theme} from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TablePagination from '@material-ui/core/TablePagination';
import TableRow from '@material-ui/core/TableRow';
import TableSortLabel from '@material-ui/core/TableSortLabel';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Paper from '@material-ui/core/Paper';
import IconButton from '@material-ui/core/IconButton';
import Tooltip from '@material-ui/core/Tooltip';
import DeleteIcon from '@material-ui/icons/Delete';
import FilterListIcon from '@material-ui/icons/FilterList';

interface Data {
  id: number;
  transaction_date: string;
  effective_date: string;
  description: string;
  category: string;
  amount: string;
}

function createData(
  id: number,
  transaction_date: string,
  effective_date: string,
  description: string,
  category: string,
  amount: string,
): Data {
  return {id, transaction_date, effective_date, description, category, amount};
}

const rows = [
  createData(1, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(2, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(3, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(4, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(5, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(6, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(7, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(8, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
  createData(9, '24/10/2020', '24/10/2020', 'Mc Donalds drive thru', 'Entertainment', '€ 340.00'),
];

function descendingComparator<T>(a: T, b: T, orderBy: keyof T) {
  if (b[orderBy] < a[orderBy]) {
    return -1;
  }
  if (b[orderBy] > a[orderBy]) {
    return 1;
  }
  return 0;
}

type Order = 'asc' | 'desc';

function getComparator<Key extends keyof any>(
  order: Order,
  orderBy: Key,
): (a: { [key in Key]: number | string }, b: { [key in Key]: number | string }) => number {
  return order === 'desc'
    ? (a, b) => descendingComparator(a, b, orderBy)
    : (a, b) => -descendingComparator(a, b, orderBy);
}

function stableSort<T>(array: T[], comparator: (a: T, b: T) => number) {
  const stabilizedThis = array.map((el, index) => [el, index] as [T, number]);
  stabilizedThis.sort((a, b) => {
    const order = comparator(a[0], b[0]);
    if (order !== 0) return order;
    return a[1] - b[1];
  });
  return stabilizedThis.map((el) => el[0]);
}

interface HeadCell {
  disablePadding: boolean;
  id: keyof Data;
  label: string;
  numeric: boolean;
}

const headCells: HeadCell[] = [
  {id: 'transaction_date', numeric: true, disablePadding: false, label: 'Transaction Date'},
  {id: 'effective_date', numeric: true, disablePadding: false, label: 'Effective Date'},
  {id: 'description', numeric: true, disablePadding: false, label: 'Description)'},
  {id: 'category', numeric: true, disablePadding: false, label: 'Category'},
  {id: 'amount', numeric: true, disablePadding: false, label: 'Amount'},
];

interface EnhancedTableProps {
  classes: ReturnType<typeof useStyles>;
  numSelected: number;
  onRequestSort: (event: React.MouseEvent<unknown>, property: keyof Data) => void;
  onSelectAllClick: (event: React.ChangeEvent<HTMLInputElement>) => void;
  order: Order;
  orderBy: string;
  rowCount: number;
}

function EnhancedTableHead(props: EnhancedTableProps) {
  const {classes, onSelectAllClick, order, orderBy, numSelected, rowCount, onRequestSort} = props;
  const createSortHandler = (property: keyof Data) => (event: React.MouseEvent<unknown>) => {
    onRequestSort(event, property);
  };

  return (
    <TableHead>
      <TableRow className={'analytics-table-head'}>
        {/*<TableCell padding="checkbox">*/}
        {/*  <Checkbox*/}
        {/*    indeterminate={numSelected > 0 && numSelected < rowCount}*/}
        {/*    checked={rowCount > 0 && numSelected === rowCount}*/}
        {/*    onChange={onSelectAllClick}*/}
        {/*    inputProps={{'aria-label': 'select all desserts'}}*/}
        {/*  />*/}
        {/*</TableCell>*/}
        {headCells.map((headCell) => (
          <TableCell
            key={headCell.id}
            align={'left'}
            padding={headCell.disablePadding ? 'none' : 'default'}
            sortDirection={orderBy === headCell.id ? order : false}
          >
            <TableSortLabel
              active={orderBy === headCell.id}
              direction={orderBy === headCell.id ? order : 'asc'}
              onClick={createSortHandler(headCell.id)}
            >
              {headCell.label}
              {orderBy === headCell.id ? (
                <span className={classes.visuallyHidden}>
                  {order === 'desc' ? 'sorted descending' : 'sorted ascending'}
                </span>
              ) : null}
            </TableSortLabel>
          </TableCell>
        ))}
      </TableRow>
    </TableHead>
  );
}

const useToolbarStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      paddingLeft: theme.spacing(2),
      paddingRight: theme.spacing(1),
    },
    highlight:
      theme.palette.type === 'light'
        ? {
          color: theme.palette.primary.main,
          backgroundColor: lighten(theme.palette.primary.light, 0.85),
        }
        : {
          color: theme.palette.text.primary,
          backgroundColor: theme.palette.primary.dark,
        },
    title: {
      flex: '1 1 100%',
    },
  }),
);

interface EnhancedTableToolbarProps {
  numSelected: number;
}

const EnhancedTableToolbar = (props: EnhancedTableToolbarProps) => {
  const classes = useToolbarStyles();
  const {numSelected} = props;

  return (
    <Toolbar
      className={clsx(classes.root, {
        [classes.highlight]: numSelected > 0,
      })}
    >
      {numSelected > 0 ? (
        <Typography className={classes.title} color="inherit" variant="subtitle1" component="div">
          {numSelected} selected
        </Typography>
      ) : (
        <Typography className={classes.title} variant="h6" id="tableTitle" component="div">
          Nutrition
        </Typography>
      )}
      {numSelected > 0 ? (
        <div/>
        // <Tooltip title="Delete">
        //   <IconButton aria-label="delete">
        //     <DeleteIcon/>
        //   </IconButton>
        // </Tooltip>
      ) : (
        <Tooltip title="Filter list">
          <IconButton aria-label="filter list">
            <FilterListIcon/>
          </IconButton>
        </Tooltip>
      )}
    </Toolbar>
  );
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      width: '100%',
    },
    paper: {
      width: '100%',
      height: '100%',
      position: 'relative'
    },
    table: {
      //minWidth: 750,
    },
    visuallyHidden: {
      border: 0,
      clip: 'rect(0 0 0 0)',
      height: 1,
      margin: -1,
      overflow: 'hidden',
      padding: 0,
      position: 'absolute',
      top: 20,
      width: 1,
    },
    tablePagination: {
      position: 'absolute',
      bottom: 0,
      right: 0
    }
  }),
);

export default function AnalyticsTable({style = {}}) {
  const classes = useStyles();
  const [order, setOrder] = React.useState<Order>('asc');
  const [orderBy, setOrderBy] = React.useState<keyof Data>('transaction_date');
  const [selected, setSelected] = React.useState<number[]>([]);
  const [page, setPage] = React.useState(0);
  const [dense, setDense] = React.useState(false);
  const [rowsPerPage, setRowsPerPage] = React.useState(5);

  useEffect(() => {
    const rootHeight = document.getElementById('analytics-table-root')?.clientHeight;
    const headHeight = document.getElementsByClassName('analytics-table-head')?.item(0)?.clientHeight;
    const rowHeight = document.getElementsByClassName('analytics-table-row')?.item(0)?.clientHeight;
    const paginationHeight = document.getElementsByClassName('analytics-table-pagination')?.item(0)?.clientHeight;

    if (rootHeight && headHeight && rowHeight && paginationHeight) {
      setRowsPerPage(Math.floor((rootHeight - headHeight - paginationHeight) / rowHeight))
    }
  }, []);

  const handleRequestSort = (event: React.MouseEvent<unknown>, property: keyof Data) => {
    const isAsc = orderBy === property && order === 'asc';
    setOrder(isAsc ? 'desc' : 'asc');
    setOrderBy(property);
  };

  const handleSelectAllClick = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.checked) {
      const newSelecteds = rows.map((n) => n.id);
      setSelected(newSelecteds);
      return;
    }
    setSelected([]);
  };

  const handleClick = (event: React.MouseEvent<unknown>, id: number) => {
    const selectedIndex = selected.indexOf(id);
    let newSelected: number[] = [];

    if (selectedIndex === -1) {
      newSelected = newSelected.concat(selected, id);
    } else if (selectedIndex === 0) {
      newSelected = newSelected.concat(selected.slice(1));
    } else if (selectedIndex === selected.length - 1) {
      newSelected = newSelected.concat(selected.slice(0, -1));
    } else if (selectedIndex > 0) {
      newSelected = newSelected.concat(
        selected.slice(0, selectedIndex),
        selected.slice(selectedIndex + 1),
      );
    }

    setSelected(newSelected);
  };

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (event: React.ChangeEvent<HTMLInputElement>) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const handleChangeDense = (event: React.ChangeEvent<HTMLInputElement>) => {
    setDense(event.target.checked);
  };

  const isSelected = (id: number) => selected.indexOf(id) !== -1;

  const emptyRows = rowsPerPage - Math.min(rowsPerPage, rows.length - page * rowsPerPage);

  return (
    <div className={classes.root} id="analytics-table-root" style={style}>
      <Paper className={classes.paper}>
        {/*<EnhancedTableToolbar numSelected={selected.length}/>*/}
        <TableContainer>
          <Table
            className={classes.table}
            aria-labelledby="tableTitle"
            size={dense ? 'small' : 'medium'}
            aria-label="enhanced table"
          >
            <EnhancedTableHead
              classes={classes}
              numSelected={selected.length}
              order={order}
              orderBy={orderBy}
              onSelectAllClick={handleSelectAllClick}
              onRequestSort={handleRequestSort}
              rowCount={rows.length}
            />
            <TableBody>
              {stableSort(rows, getComparator(order, orderBy))
                .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
                .map((row, index) => {
                  const isItemSelected = isSelected(row.id);
                  const labelId = `enhanced-analytics-table-checkbox-${index}`;

                  return (
                    <TableRow
                      hover
                      className={'analytics-table-row'}
                      onClick={(event) => handleClick(event, row.id)}
                      role="checkbox"
                      aria-checked={isItemSelected}
                      tabIndex={-1}
                      key={row.id}
                      selected={isItemSelected}
                    >
                      {/*<TableCell padding="checkbox">*/}
                      {/*  <Checkbox*/}
                      {/*    checked={isItemSelected}*/}
                      {/*    inputProps={{'aria-labelledby': labelId}}*/}
                      {/*  />*/}
                      {/*</TableCell>*/}
                      <TableCell id={labelId} scope="row" align="left">
                        <span style={{background: '#ECECEC', padding: 4}}>{row.transaction_date}</span>
                      </TableCell>
                      <TableCell align="left">
                        <span style={{background: '#ECECEC', padding: 4}}>{row.effective_date}</span>
                      </TableCell>
                      <TableCell align="left">{row.description}</TableCell>
                      <TableCell align="left">
                        <div>
                          <div style={{display: 'inline-block', width: 12, height: 12, borderRadius: '50%', background: '#3098D6', position: 'relative', top: 1, marginRight: 4}}/>
                          {row.category}
                        </div>
                      </TableCell>
                      <TableCell align="left">{row.amount}</TableCell>
                    </TableRow>
                  );
                })}
              {emptyRows > 0 && (
                <TableRow style={{height: (dense ? 33 : 53) * emptyRows}}>
                  <TableCell colSpan={6}/>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </TableContainer>
        <TablePagination
          className={clsx(['analytics-table-pagination', classes.tablePagination])}
          rowsPerPageOptions={[]}
          component="div"
          count={rows.length}
          rowsPerPage={rowsPerPage}
          page={page}
          onChangePage={handleChangePage}
          onChangeRowsPerPage={handleChangeRowsPerPage}
        />
      </Paper>
    </div>
  );
}
