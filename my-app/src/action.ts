/*
 * action types
 */

export const CLICK_SQUARE = "CLICK_SQUARE";

/*
  * action creators
  */

export function clickSquare(index: any) {
    return {type: CLICK_SQUARE, index};
}