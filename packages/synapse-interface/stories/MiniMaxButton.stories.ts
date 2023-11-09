import type { Meta, StoryObj } from '@storybook/react'

import MiniMaxButton from '@/components/buttons/MiniMaxButton'
// More on how to set up stories at: https://storybook.js.org/docs/react/writing-stories/introduction#default-export
const meta = {
  title: 'Synapse/MiniMaxButton',
  component: MiniMaxButton,
  parameters: {
    // Optional parameter to center the component in the Canvas. More info: https://storybook.js.org/docs/react/configure/story-layout
    layout: 'centered',
  },
  // This component will have an automatically generated Autodocs entry: https://storybook.js.org/docs/react/writing-docs/autodocs
  tags: ['autodocs'],
  // More on argTypes: https://storybook.js.org/docs/react/api/argtypes
  argTypes: {
    onClickBalance: () => void 0,
    disabled: { control: 'boolean' },
  },
} satisfies Meta<typeof MiniMaxButton>

export default meta
type Story = StoryObj<typeof meta>

// More on writing stories with args: https://storybook.js.org/docs/react/writing-stories/args
export const Primary: Story = {
  args: {
    onClickBalance: () => void 0,
    disabled: false,
  },
}
