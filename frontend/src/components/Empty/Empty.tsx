import {
  Link
} from "react-router-dom";

const Empty = (): JSX.Element => {
  return (
    <div>
      <h2>There are no bikes to show</h2>
      <p>What about adding some??</p>

      <Link to="/add">Add bike</Link>
    </div>
  )
};

export default Empty;
