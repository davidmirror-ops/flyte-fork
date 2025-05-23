package update

import (
	"fmt"
	"testing"

	"github.com/flyteorg/flyte/flytectl/cmd/config/subcommand/executionqueueattribute"
	"github.com/flyteorg/flyte/flytectl/cmd/testutils"
	"github.com/flyteorg/flyte/flytectl/pkg/ext"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	validWorkflowExecutionQueueMatchableAttributesFilePath     = "testdata/valid_workflow_execution_queue_attribute.yaml"
	validProjectDomainExecutionQueueMatchableAttributeFilePath = "testdata/valid_project_domain_execution_queue_attribute.yaml"
	validProjectExecutionQueueMatchableAttributeFilePath       = "testdata/valid_project_execution_queue_attribute.yaml"
)

func TestExecutionQueueAttributeUpdateRequiresAttributeFile(t *testing.T) {
	testWorkflowExecutionQueueAttributeUpdate(
		t,
		/* setup */ nil,
		/* assert */ func(s *testutils.TestStruct, err error) {
			assert.ErrorContains(t, err, "attrFile is mandatory")
			s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
		})
}

func TestExecutionQueueAttributeUpdateFailsWhenAttributeFileDoesNotExist(t *testing.T) {
	testWorkflowExecutionQueueAttributeUpdate(
		t,
		/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
			config.AttrFile = testDataNonExistentFile
			config.Force = true
		},
		/* assert */ func(s *testutils.TestStruct, err error) {
			assert.ErrorContains(t, err, "unable to read from testdata/non-existent-file yaml file")
			s.UpdaterExt.AssertNotCalled(t, "FetchWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
		})
}

func TestExecutionQueueAttributeUpdateFailsWhenAttributeFileIsMalformed(t *testing.T) {
	testWorkflowExecutionQueueAttributeUpdate(
		t,
		/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
			config.AttrFile = testDataInvalidAttrFile
			config.Force = true
		},
		/* assert */ func(s *testutils.TestStruct, err error) {
			assert.ErrorContains(t, err, "error unmarshaling JSON: while decoding JSON: json: unknown field \"InvalidDomain\"")
			s.UpdaterExt.AssertNotCalled(t, "FetchWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
		})
}

func TestExecutionQueueAttributeUpdateHappyPath(t *testing.T) {
	t.Run("workflow", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				s.TearDownAndVerifyContains(t, `Updated attributes from flytesnacks project and domain development`)
			})
	})

	t.Run("domain", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				s.TearDownAndVerifyContains(t, `Updated attributes from flytesnacks project and domain development`)
			})
	})

	t.Run("project", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
				s.TearDownAndVerifyContains(t, `Updated attributes from flytesnacks project`)
			})
	})
}

func TestExecutionQueueAttributeUpdateFailsWithoutForceFlag(t *testing.T) {
	t.Run("workflow", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.Force = false
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.ErrorContains(t, err, "update aborted by user")
				s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("domain", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.Force = false
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.ErrorContains(t, err, "update aborted by user")
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("project", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.Force = false
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.ErrorContains(t, err, "update aborted by user")
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
			})
	})
}

func TestExecutionQueueAttributeUpdateDoesNothingWithDryRunFlag(t *testing.T) {
	t.Run("workflow", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("domain", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("project", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
			})
	})
}

func TestExecutionQueueAttributeUpdateIgnoresForceFlagWithDryRun(t *testing.T) {
	t.Run("workflow without --force", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.Force = false
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("workflow with --force", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.Force = true
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("domain without --force", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.Force = false
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("domain with --force", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.Force = true
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("project without --force", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.Force = false
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("project with --force", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdate(
			t,
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.Force = true
				config.DryRun = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertNotCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
			})
	})
}

func TestExecutionQueueAttributeUpdateSucceedsWhenAttributesDoNotExist(t *testing.T) {
	t.Run("workflow", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdateWithMockSetup(
			t,
			/* mockSetup */ func(s *testutils.TestStruct, target *admin.WorkflowAttributes) {
				s.FetcherExt.
					EXPECT().FetchWorkflowAttributes(s.Ctx, target.GetProject(), target.GetDomain(), target.GetWorkflow(), admin.MatchableResource_EXECUTION_QUEUE).
					Return(nil, ext.NewNotFoundError("attribute"))
				s.UpdaterExt.
					EXPECT().UpdateWorkflowAttributes(s.Ctx, target.GetProject(), target.GetDomain(), target.GetWorkflow(), mock.Anything).
					Return(nil)
			},
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				s.TearDownAndVerifyContains(t, `Updated attributes from flytesnacks project and domain development`)
			})
	})

	t.Run("domain", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdateWithMockSetup(
			t,
			/* mockSetup */ func(s *testutils.TestStruct, target *admin.ProjectDomainAttributes) {
				s.FetcherExt.
					EXPECT().FetchProjectDomainAttributes(s.Ctx, target.GetProject(), target.GetDomain(), admin.MatchableResource_EXECUTION_QUEUE).
					Return(nil, ext.NewNotFoundError("attribute"))
				s.UpdaterExt.
					EXPECT().UpdateProjectDomainAttributes(s.Ctx, target.GetProject(), target.GetDomain(), mock.Anything).
					Return(nil)
			},
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
				s.TearDownAndVerifyContains(t, `Updated attributes from flytesnacks project and domain development`)
			})
	})

	t.Run("project", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdateWithMockSetup(
			t,
			/* mockSetup */ func(s *testutils.TestStruct, target *admin.ProjectAttributes) {
				s.FetcherExt.
					EXPECT().FetchProjectAttributes(s.Ctx, target.GetProject(), admin.MatchableResource_EXECUTION_QUEUE).
					Return(nil, ext.NewNotFoundError("attribute"))
				s.UpdaterExt.
					EXPECT().UpdateProjectAttributes(s.Ctx, target.GetProject(), mock.Anything).
					Return(nil)
			},
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Nil(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
				s.TearDownAndVerifyContains(t, `Updated attributes from flytesnacks project`)
			})
	})
}

func TestExecutionQueueAttributeUpdateFailsWhenAdminClientFails(t *testing.T) {
	t.Run("workflow", func(t *testing.T) {
		testWorkflowExecutionQueueAttributeUpdateWithMockSetup(
			t,
			/* mockSetup */ func(s *testutils.TestStruct, target *admin.WorkflowAttributes) {
				s.FetcherExt.
					EXPECT().FetchWorkflowAttributes(s.Ctx, target.GetProject(), target.GetDomain(), target.GetWorkflow(), admin.MatchableResource_EXECUTION_QUEUE).
					Return(&admin.WorkflowAttributesGetResponse{Attributes: target}, nil)
				s.UpdaterExt.
					EXPECT().UpdateWorkflowAttributes(s.Ctx, target.GetProject(), target.GetDomain(), target.GetWorkflow(), mock.Anything).
					Return(fmt.Errorf("network error"))
			},
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes) {
				config.AttrFile = validWorkflowExecutionQueueMatchableAttributesFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Error(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateWorkflowAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("domain", func(t *testing.T) {
		testProjectDomainExecutionQueueAttributeUpdateWithMockSetup(
			t,
			/* mockSetup */ func(s *testutils.TestStruct, target *admin.ProjectDomainAttributes) {
				s.FetcherExt.
					EXPECT().FetchProjectDomainAttributes(s.Ctx, target.GetProject(), target.GetDomain(), admin.MatchableResource_EXECUTION_QUEUE).
					Return(&admin.ProjectDomainAttributesGetResponse{Attributes: target}, nil)
				s.UpdaterExt.
					EXPECT().UpdateProjectDomainAttributes(s.Ctx, target.GetProject(), target.GetDomain(), mock.Anything).
					Return(fmt.Errorf("network error"))
			},
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes) {
				config.AttrFile = validProjectDomainExecutionQueueMatchableAttributeFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Error(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateProjectDomainAttributes", mock.Anything, mock.Anything, mock.Anything, mock.Anything)
			})
	})

	t.Run("project", func(t *testing.T) {
		testProjectExecutionQueueAttributeUpdateWithMockSetup(
			t,
			/* mockSetup */ func(s *testutils.TestStruct, target *admin.ProjectAttributes) {
				s.FetcherExt.
					EXPECT().FetchProjectAttributes(s.Ctx, target.GetProject(), admin.MatchableResource_EXECUTION_QUEUE).
					Return(&admin.ProjectAttributesGetResponse{Attributes: target}, nil)
				s.UpdaterExt.
					EXPECT().UpdateProjectAttributes(s.Ctx, target.GetProject(), mock.Anything).
					Return(fmt.Errorf("network error"))
			},
			/* setup */ func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes) {
				config.AttrFile = validProjectExecutionQueueMatchableAttributeFilePath
				config.Force = true
			},
			/* assert */ func(s *testutils.TestStruct, err error) {
				assert.Error(t, err)
				s.UpdaterExt.AssertCalled(t, "UpdateProjectAttributes", mock.Anything, mock.Anything, mock.Anything)
			})
	})
}

func testWorkflowExecutionQueueAttributeUpdate(
	t *testing.T,
	setup func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes),
	asserter func(s *testutils.TestStruct, err error),
) {
	testWorkflowExecutionQueueAttributeUpdateWithMockSetup(
		t,
		/* mockSetup */ func(s *testutils.TestStruct, target *admin.WorkflowAttributes) {
			s.FetcherExt.
				EXPECT().FetchWorkflowAttributes(s.Ctx, target.GetProject(), target.GetDomain(), target.GetWorkflow(), admin.MatchableResource_EXECUTION_QUEUE).
				Return(&admin.WorkflowAttributesGetResponse{Attributes: target}, nil)
			s.UpdaterExt.
				EXPECT().UpdateWorkflowAttributes(s.Ctx, target.GetProject(), target.GetDomain(), target.GetWorkflow(), mock.Anything).
				Return(nil)
		},
		setup,
		asserter,
	)
}

func testWorkflowExecutionQueueAttributeUpdateWithMockSetup(
	t *testing.T,
	mockSetup func(s *testutils.TestStruct, target *admin.WorkflowAttributes),
	setup func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.WorkflowAttributes),
	asserter func(s *testutils.TestStruct, err error),
) {
	s := testutils.Setup(t)

	executionqueueattribute.DefaultUpdateConfig = &executionqueueattribute.AttrUpdateConfig{}
	target := newTestWorkflowExecutionQueueAttribute()

	if mockSetup != nil {
		mockSetup(&s, target)
	}

	if setup != nil {
		setup(&s, executionqueueattribute.DefaultUpdateConfig, target)
	}

	err := updateExecutionQueueAttributesFunc(s.Ctx, nil, s.CmdCtx)

	if asserter != nil {
		asserter(&s, err)
	}

	// cleanup
	executionqueueattribute.DefaultUpdateConfig = &executionqueueattribute.AttrUpdateConfig{}
}

func newTestWorkflowExecutionQueueAttribute() *admin.WorkflowAttributes {
	return &admin.WorkflowAttributes{
		// project, domain, and workflow names need to be same as in the tests spec files in testdata folder
		Project:  "flytesnacks",
		Domain:   "development",
		Workflow: "core.control_flow.merge_sort.merge_sort",
		MatchingAttributes: &admin.MatchingAttributes{
			Target: &admin.MatchingAttributes_ExecutionQueueAttributes{
				ExecutionQueueAttributes: &admin.ExecutionQueueAttributes{
					Tags: []string{
						testutils.RandomName(5),
						testutils.RandomName(5),
						testutils.RandomName(5),
					},
				},
			},
		},
	}
}

func testProjectExecutionQueueAttributeUpdate(
	t *testing.T,
	setup func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes),
	asserter func(s *testutils.TestStruct, err error),
) {
	testProjectExecutionQueueAttributeUpdateWithMockSetup(
		t,
		/* mockSetup */ func(s *testutils.TestStruct, target *admin.ProjectAttributes) {
			s.FetcherExt.
				EXPECT().FetchProjectAttributes(s.Ctx, target.GetProject(), admin.MatchableResource_EXECUTION_QUEUE).
				Return(&admin.ProjectAttributesGetResponse{Attributes: target}, nil)
			s.UpdaterExt.
				EXPECT().UpdateProjectAttributes(s.Ctx, target.GetProject(), mock.Anything).
				Return(nil)
		},
		setup,
		asserter,
	)
}

func testProjectExecutionQueueAttributeUpdateWithMockSetup(
	t *testing.T,
	mockSetup func(s *testutils.TestStruct, target *admin.ProjectAttributes),
	setup func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectAttributes),
	asserter func(s *testutils.TestStruct, err error),
) {
	s := testutils.Setup(t)

	executionqueueattribute.DefaultUpdateConfig = &executionqueueattribute.AttrUpdateConfig{}
	target := newTestProjectExecutionQueueAttribute()

	if mockSetup != nil {
		mockSetup(&s, target)
	}

	if setup != nil {
		setup(&s, executionqueueattribute.DefaultUpdateConfig, target)
	}

	err := updateExecutionQueueAttributesFunc(s.Ctx, nil, s.CmdCtx)

	if asserter != nil {
		asserter(&s, err)
	}

	// cleanup
	executionqueueattribute.DefaultUpdateConfig = &executionqueueattribute.AttrUpdateConfig{}
}

func newTestProjectExecutionQueueAttribute() *admin.ProjectAttributes {
	return &admin.ProjectAttributes{
		// project name needs to be same as in the tests spec files in testdata folder
		Project: "flytesnacks",
		MatchingAttributes: &admin.MatchingAttributes{
			Target: &admin.MatchingAttributes_ExecutionQueueAttributes{
				ExecutionQueueAttributes: &admin.ExecutionQueueAttributes{
					Tags: []string{
						testutils.RandomName(5),
						testutils.RandomName(5),
						testutils.RandomName(5),
					},
				},
			},
		},
	}
}

func testProjectDomainExecutionQueueAttributeUpdate(
	t *testing.T,
	setup func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes),
	asserter func(s *testutils.TestStruct, err error),
) {
	testProjectDomainExecutionQueueAttributeUpdateWithMockSetup(
		t,
		/* mockSetup */ func(s *testutils.TestStruct, target *admin.ProjectDomainAttributes) {
			s.FetcherExt.
				EXPECT().FetchProjectDomainAttributes(s.Ctx, target.GetProject(), target.GetDomain(), admin.MatchableResource_EXECUTION_QUEUE).
				Return(&admin.ProjectDomainAttributesGetResponse{Attributes: target}, nil)
			s.UpdaterExt.
				EXPECT().UpdateProjectDomainAttributes(s.Ctx, target.GetProject(), target.GetDomain(), mock.Anything).
				Return(nil)
		},
		setup,
		asserter,
	)
}

func testProjectDomainExecutionQueueAttributeUpdateWithMockSetup(
	t *testing.T,
	mockSetup func(s *testutils.TestStruct, target *admin.ProjectDomainAttributes),
	setup func(s *testutils.TestStruct, config *executionqueueattribute.AttrUpdateConfig, target *admin.ProjectDomainAttributes),
	asserter func(s *testutils.TestStruct, err error),
) {
	s := testutils.Setup(t)

	executionqueueattribute.DefaultUpdateConfig = &executionqueueattribute.AttrUpdateConfig{}
	target := newTestProjectDomainExecutionQueueAttribute()

	if mockSetup != nil {
		mockSetup(&s, target)
	}

	if setup != nil {
		setup(&s, executionqueueattribute.DefaultUpdateConfig, target)
	}

	err := updateExecutionQueueAttributesFunc(s.Ctx, nil, s.CmdCtx)

	if asserter != nil {
		asserter(&s, err)
	}

	// cleanup
	executionqueueattribute.DefaultUpdateConfig = &executionqueueattribute.AttrUpdateConfig{}
}

func newTestProjectDomainExecutionQueueAttribute() *admin.ProjectDomainAttributes {
	return &admin.ProjectDomainAttributes{
		// project and domain names need to be same as in the tests spec files in testdata folder
		Project: "flytesnacks",
		Domain:  "development",
		MatchingAttributes: &admin.MatchingAttributes{
			Target: &admin.MatchingAttributes_ExecutionQueueAttributes{
				ExecutionQueueAttributes: &admin.ExecutionQueueAttributes{
					Tags: []string{
						testutils.RandomName(5),
						testutils.RandomName(5),
						testutils.RandomName(5),
					},
				},
			},
		},
	}
}
