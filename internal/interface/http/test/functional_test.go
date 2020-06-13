package test

import (
	"dast-api/internal/interface/http/presenter"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestComplexExample(t *testing.T) {
	assert := assert.New(t)
	resCreateHierarchy := performRequest("POST", "/hierarchy", `{"name": "testing complex case", "description": "esta es una jerarquía de prueba", "owner": "guest", "objective": "Elegir auto", "alternatives": ["Ford", "Fiat", "Chevrolet"] }`, controllers)
	assert.Equal(http.StatusCreated, resCreateHierarchy.Code)

	var hierarchyResponse presenter.HierarchyResponse
	json.Unmarshal(resCreateHierarchy.Body.Bytes(), &hierarchyResponse)
	assert.NotNil(hierarchyResponse.ID)

	resFillCriteria := performRequest("PUT", fmt.Sprintf("/hierarchy/%s/criteria", hierarchyResponse.ID), `[{"level": 0, "description": "Velocidad", "id": "velocidad"}, {"level": 1, "description": "Velocidad 1", "id": "velocidad_1", "parent": "velocidad"}, {"level": 1, "description": "Velocidad 2", "id": "velocidad_2", "parent": "velocidad"}, {"level": 0, "description": "Cilindrada", "id": "cilindrada"}, {"level": 0, "description": "Aceleración", "id": "aceleracion"}, {"level": 1, "description": "Aceleracion 1", "id": "aceleracion_1", "parent": "aceleracion"}, {"level": 1, "description": "Aceleracion 2", "id": "aceleracion_2", "parent": "aceleracion"}, {"level": 2, "description": "Aceleracion 21", "id": "aceleracion_21", "parent": "aceleracion_2"}, {"level": 2, "description": "Aceleracion 22", "id": "aceleracion_22", "parent": "aceleracion_2"} ]`, controllers)
	assert.Equal(http.StatusOK, resFillCriteria.Code)

	get := performRequest("GET", fmt.Sprintf("/hierarchy/%s", hierarchyResponse.ID), ``, controllers)
	assert.Equal(http.StatusOK, get.Code)

	var h presenter.HierarchyResponse
	json.Unmarshal(get.Body.Bytes(), &h)
	assert.NotNil(h.Criteria)
	assert.Equal(9, len(h.Criteria))

	resPWiseGenerationHTTP := performRequest("POST", fmt.Sprintf("/pairwise/%s/generate", h.ID), ``, controllers)
	assert.Equal(http.StatusCreated, resPWiseGenerationHTTP.Code)

	var pWise presenter.CriteriaJudgements
	json.Unmarshal(resPWiseGenerationHTTP.Body.Bytes(), &pWise)

	assert.NotNil(pWise.ID)
	assert.Equal(4, len(pWise.CriteriaComparison))
	assert.Equal(6, len(pWise.AlternativeComparison))

	resPWiseFillHTTP := performRequest("PUT", fmt.Sprintf("/pairwise/%s/judgements/%s", pWise.HierarchyID, pWise.ID), `{"criteria_comparison": [{"level": 0, "matrix_context": {"compared_to": "", "elements": ["velocidad", "cilindrada", "aceleracion"], "judgements": [[1, 2, 0.5 ], [0.5, 1, 3 ], [2, 0.3333333333333, 1 ] ] } }, {"level": 1, "matrix_context": {"compared_to": "velocidad", "elements": ["velocidad_1", "velocidad_2"], "judgements": [[1, 2 ], [0.5, 1 ] ] } }, {"level": 1, "matrix_context": {"compared_to": "aceleracion", "elements": ["aceleracion_1", "aceleracion_2"], "judgements": [[1, 5 ], [0.2, 1 ] ] } }, {"level": 2, "matrix_context": {"compared_to": "aceleracion_2", "elements": ["aceleracion_21", "aceleracion_22"], "judgements": [[1, 0.4 ], [2.5, 1 ] ] } } ], "alternative_comparison": [{"compared_to": "velocidad_1", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 0.5, 3 ], [2, 1, 0.5 ], [0.333333333333, 2, 1 ] ] }, {"compared_to": "velocidad_2", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 3 ], [0.5, 1, 4 ], [0.333333333333, 0.25, 1 ] ] }, {"compared_to": "cilindrada", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 0.5 ], [0.5, 1, 0.25 ], [2, 4, 1 ] ] }, {"compared_to": "aceleracion_1", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 0.5 ], [0.5, 1, 0.33333333333333 ], [2, 3, 1 ] ] }, {"compared_to": "aceleracion_21", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 4, 5 ], [0.25, 1, 3 ], [0.2, 0.3333333333333, 1 ] ] }, {"compared_to": "aceleracion_22", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 1 ], [0.5, 1, 0.3333333333333 ], [1, 3, 1 ] ] } ] }`, controllers)
	assert.Equal(http.StatusOK, resPWiseFillHTTP.Code)

	resResolveHTTP := performRequest("POST", fmt.Sprintf("/pairwise/%s/judgements/%s/resolve", pWise.HierarchyID, pWise.ID), ``, controllers)
	assert.Equal(http.StatusOK, resResolveHTTP.Code)

	var resolve presenter.CriteriaJudgements
	json.Unmarshal(resResolveHTTP.Body.Bytes(), &resolve)

	assert.Equal(0.34125622741095857, resolve.Results["Ford"])
	assert.Equal(0.2162846496027526, resolve.Results["Fiat"])
	assert.Equal(0.4424591229862888, resolve.Results["Chevrolet"])
}

func TestDisorderComplexExample(t *testing.T) {
	assert := assert.New(t)
	resCreateHierarchy := performRequest("POST", "/hierarchy", `{"name": "testing complex case", "description": "esta es una jerarquía de prueba", "owner": "amarini", "objective": "Elegir auto", "alternatives": ["Ford", "Fiat", "Chevrolet"] }`, controllers)
	assert.Equal(http.StatusCreated, resCreateHierarchy.Code)

	var hierarchyResponse presenter.HierarchyResponse
	json.Unmarshal(resCreateHierarchy.Body.Bytes(), &hierarchyResponse)
	assert.NotNil(hierarchyResponse.ID)

	resFillCriteria := performRequest("PUT", fmt.Sprintf("/hierarchy/%s/criteria", hierarchyResponse.ID), `[{"level": 0, "description": "Velocidad", "id": "velocidad"}, {"level": 1, "description": "Velocidad 1", "id": "velocidad_1", "parent": "velocidad"}, {"level": 1, "description": "Velocidad 2", "id": "velocidad_2", "parent": "velocidad"}, {"level": 0, "description": "Cilindrada", "id": "cilindrada"}, {"level": 0, "description": "Aceleración", "id": "aceleracion"}, {"level": 1, "description": "Aceleracion 1", "id": "aceleracion_1", "parent": "aceleracion"}, {"level": 1, "description": "Aceleracion 2", "id": "aceleracion_2", "parent": "aceleracion"}, {"level": 2, "description": "Aceleracion 21", "id": "aceleracion_21", "parent": "aceleracion_2"}, {"level": 2, "description": "Aceleracion 22", "id": "aceleracion_22", "parent": "aceleracion_2"} ]`, controllers)
	assert.Equal(http.StatusOK, resFillCriteria.Code)

	get := performRequest("GET", fmt.Sprintf("/hierarchy/%s", hierarchyResponse.ID), ``, controllers)
	assert.Equal(http.StatusOK, get.Code)

	var h presenter.HierarchyResponse
	json.Unmarshal(get.Body.Bytes(), &h)
	assert.NotNil(h.Criteria)
	assert.Equal(9, len(h.Criteria))

	resPWiseGenerationHTTP := performRequest("POST", fmt.Sprintf("/pairwise/%s/generate", h.ID), ``, controllers)
	assert.Equal(http.StatusCreated, resPWiseGenerationHTTP.Code)

	var pWise presenter.CriteriaJudgements
	json.Unmarshal(resPWiseGenerationHTTP.Body.Bytes(), &pWise)

	assert.NotNil(pWise.ID)
	assert.Equal(4, len(pWise.CriteriaComparison))
	assert.Equal(6, len(pWise.AlternativeComparison))

	resPWiseFillHTTP := performRequest("PUT", fmt.Sprintf("/pairwise/%s/judgements/%s", pWise.HierarchyID, pWise.ID), `{"criteria_comparison": [{"level": 1, "matrix_context": {"compared_to": "velocidad", "elements": ["velocidad_1", "velocidad_2"], "judgements": [[1, 2 ], [0.5, 1 ] ] } }, {"level": 1, "matrix_context": {"compared_to": "aceleracion", "elements": ["aceleracion_1", "aceleracion_2"], "judgements": [[1, 5 ], [0.2, 1 ] ] } }, {"level": 2, "matrix_context": {"compared_to": "aceleracion_2", "elements": ["aceleracion_21", "aceleracion_22"], "judgements": [[1, 0.4 ], [2.5, 1 ] ] } }, {"level": 0, "matrix_context": {"compared_to": "", "elements": ["velocidad", "cilindrada", "aceleracion"], "judgements": [[1, 2, 0.5 ], [0.5, 1, 3 ], [2, 0.3333333333333, 1 ] ] } } ], "alternative_comparison": [{"compared_to": "velocidad_2", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 3 ], [0.5, 1, 4 ], [0.333333333333, 0.25, 1 ] ] }, {"compared_to": "cilindrada", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 0.5 ], [0.5, 1, 0.25 ], [2, 4, 1 ] ] }, {"compared_to": "aceleracion_1", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 0.5 ], [0.5, 1, 0.33333333333333 ], [2, 3, 1 ] ] }, {"compared_to": "aceleracion_21", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 4, 5 ], [0.25, 1, 3 ], [0.2, 0.3333333333333, 1 ] ] }, {"compared_to": "aceleracion_22", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 2, 1 ], [0.5, 1, 0.3333333333333 ], [1, 3, 1 ] ] }, {"compared_to": "velocidad_1", "elements": ["Ford", "Fiat", "Chevrolet"], "judgements": [[1, 0.5, 3 ], [2, 1, 0.5 ], [0.333333333333, 2, 1 ] ] } ] }`, controllers)
	assert.Equal(http.StatusInternalServerError, resPWiseFillHTTP.Code)
	assert.Contains(resPWiseFillHTTP.Body.String(), "alternative comparison must be in order for ensure quality results")
}

func TestCreateAndLoginUser(t *testing.T) {
	assert := assert.New(t)
	resCreateUser := performRequest("POST", "/user/create", `{"name": "pepe", "email": "pepe@gmail.com", "password": "Atrevi2"}`, controllers)

	assert.Equal(http.StatusCreated, resCreateUser.Code)
	var user presenter.User
	json.Unmarshal(resCreateUser.Body.Bytes(), &user)
	assert.NotEmpty(user.ID)

	resLogIn := performRequest("POST", "/user/login", `{"name": "pepe", "password": "Atrevi2"}`, controllers)
	assert.Equal(http.StatusCreated, resLogIn.Code)
	var login presenter.LogResponse
	json.Unmarshal(resLogIn.Body.Bytes(), &login)
	assert.NotEmpty(login.Token)
	assert.Equal(fmt.Sprintf("%s successful logged in", "pepe"), login.Message)

	resLogIn2 := performRequest("POST", "/user/login", `{"name": "pepe", "password": "Atrevi2"}`, controllers)
	assert.Equal(http.StatusCreated, resLogIn2.Code)
	var login2 presenter.LogResponse
	json.Unmarshal(resLogIn2.Body.Bytes(), &login2)
	assert.NotEmpty(login2.Token)
	assert.Equal(login.Token, login2.Token)

	resValidateToken := performRequest("POST", "/user/validate", fmt.Sprintf(`{"id": "%s", "token": "%s"}`, user.ID, login.Token), controllers)
	assert.Equal(http.StatusNoContent, resValidateToken.Code)
}
