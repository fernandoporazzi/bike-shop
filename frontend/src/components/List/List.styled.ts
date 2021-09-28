import styled from "styled-components"

export const StyledList = styled.div`
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
`;

export const StyledListItem = styled.div`
  border-radius: 5px;
  background-color: white;
  width: 30%;
  border: 1px solid #999;
  margin-bottom: 40px;
`;

export const StyledImageWrapper = styled.div`
  width: 100%;
  height: 360px;

  img {
    max-width: 100%;
    max-height: 100%;
    border-top-left-radius: 5px;
    border-top-right-radius: 5px;
  }
`

export const StyledListItemBottom = styled.div`
  padding: 10px;
`;
