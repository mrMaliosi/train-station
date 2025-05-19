// src/layouts/Tabs.tsx
import * as RadixTabs from '@radix-ui/react-tabs';
import type { ReactNode } from 'react';

export type TabId = 'employees' | 'passengers' | 'tickets' | 'departments';
export interface Tab {
  id: TabId;
  label: string;
}

interface TabsProps {
  tabs: Tab[];
  activeTab: TabId;
  /**
   * Принимает строго TabId,
   * но обёртка в onValueChange приведёт string → TabId
   */
  onChange: (id: TabId) => void;
  children: Record<TabId, ReactNode>;
}

export default function Tabs({
  tabs,
  activeTab,
  onChange,
  children,
}: TabsProps) {
  return (
    <RadixTabs.Root
      className="space-y-4"
      value={activeTab}
      // оборачиваем, приводим string → TabId
      onValueChange={(value) => onChange(value as TabId)}
    >
      <RadixTabs.List className="tabs flex gap-2">
        {tabs.map((tab) => (
          <RadixTabs.Trigger
            key={tab.id}
            value={tab.id}
            className="px-4 py-2 rounded-lg hover:bg-gray-100 data-[state=active]:bg-blue-500 data-[state=active]:text-white"
          >
            {tab.label}
          </RadixTabs.Trigger>
        ))}
      </RadixTabs.List>

      {tabs.map((tab) => (
        <RadixTabs.Content
          key={tab.id}
          value={tab.id}
          className="p-4 border rounded-lg"
        >
          {children[tab.id]}
        </RadixTabs.Content>
      ))}
    </RadixTabs.Root>
  );
}
