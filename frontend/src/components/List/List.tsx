import { Bike } from "../../types";
import {
  Link
} from "react-router-dom";
import {
  StyledList,
  StyledListItem,
  StyledImageWrapper,
  StyledListItemBottom,
} from "./List.styled";

type Props = {
  bikes: Bike[];
};

const getBikePath = (paths: string[]): string => {
  if (paths && paths.length > 0) {
    return `http://localhost:8080/static/${paths[0]}`;
  }

  return "https://via.placeholder.com/360";
};

const List = ({ bikes }: Props): JSX.Element => {
  return (
    <StyledList>
      {bikes.map(bike => (
        <StyledListItem key={bike.id}>
          <Link to={`/bikes/${bike.id}`}>
              <StyledImageWrapper>
                <img src={getBikePath(bike.images)} alt={bike.name} />
              </StyledImageWrapper>

              <StyledListItemBottom>
                <h2>{bike.name}</h2>

                <p>Color: {bike.color}</p>
                <p>In stock: {bike.stock }</p>
              </StyledListItemBottom>
          </Link>
        </StyledListItem>
      ))}
    </StyledList>
  );
};

export default List;
