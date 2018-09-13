console.log('*** inside the store/mutations file ***');

export const testMutation = (state, {newString}) => {
    state.testString = newString;
    consoleAfterSet(state.newString)
}

export const setSingleVarMutation = (state, {variableName, variableValue}) => {
    state[variableName] = variableValue;
    consoleAfterSet(variableValue)
}

export const setMultiVarMutation = (state, {variableArray}) => {
    variableArray.forEach(variable => {
        state[variable.variableName] = variable.variableValue
        consoleAfterSet(state[variable.variableName])
        console.log('and value of state after setMultiVarMutation: ', state);
    });
}

// helpers

const consoleAfterSet = (variableName) => {
    console.log('value of variable after setting: ', variableName);
}